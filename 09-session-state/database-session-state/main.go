package main

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strconv"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/adnvilla/enterprise_patterns/09-session-state/database-session-state/sesion"
)

const ttl = 30 * time.Minute

// El cliente solo lleva un identificador opaco
func newSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func main() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "postgres://app:app@localhost:5432/tienda?sslmode=disable"
	}
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal("no puedo conectar a Postgres (¿corriste `docker compose up -d --wait`?): ", err)
	}

	ctx := context.Background()
	store := sesion.NewDBStore(db)

	// Crea la sesión: una fila más en la tabla sessions
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		customerID, _ := strconv.ParseInt(r.URL.Query().Get("customer"), 10, 64)
		id := newSessionID()
		if err := store.Save(r.Context(), id, sesion.Session{CustomerID: customerID}, ttl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.SetCookie(w, &http.Cookie{Name: "sid", Value: id, HttpOnly: true})
		fmt.Fprintf(w, "sesión creada para el cliente %d\n", customerID)
	})

	// Modifica el estado: cualquier servidor puede atender, y un reinicio no pierde nada
	http.HandleFunc("/cart/add", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sid")
		if err != nil {
			http.Error(w, "sin sesión", http.StatusUnauthorized)
			return
		}
		sess, err := store.Load(r.Context(), c.Value)
		if err != nil {
			http.Error(w, "sesión expirada o desconocida", http.StatusUnauthorized)
			return
		}
		productID, _ := strconv.ParseInt(r.URL.Query().Get("product"), 10, 64)
		sess.CartItems = append(sess.CartItems, productID)
		if err := store.Save(r.Context(), c.Value, sess, ttl); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "producto %d agregado; carrito: %v\n", productID, sess.CartItems)
	})

	// Lee el estado desde la base de datos
	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sid")
		if err != nil {
			http.Error(w, "sin sesión", http.StatusUnauthorized)
			return
		}
		sess, err := store.Load(r.Context(), c.Value)
		if err != nil {
			http.Error(w, "sesión expirada o desconocida", http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "cliente %d, carrito: %v\n", sess.CustomerID, sess.CartItems)
	})

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, nil)

	fmt.Println("Servidor escuchando en :8080")
	fmt.Println("Rutas de ejemplo:")
	fmt.Println(`  curl -c cookies.txt 'http://localhost:8080/login?customer=7'`)
	fmt.Println(`  curl -b cookies.txt 'http://localhost:8080/cart/add?product=5'`)
	fmt.Println(`  curl -b cookies.txt 'http://localhost:8080/cart'`)
	fmt.Println()

	// Demostración end-to-end contra el Postgres del compose
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	for _, url := range []string{
		"http://localhost:8080/login?customer=7",
		"http://localhost:8080/cart/add?product=1",
		"http://localhost:8080/cart/add?product=5",
		"http://localhost:8080/cart",
	} {
		resp, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("GET %-46s -> %s\n", url, resp.Status)
		resp.Body.Close()
	}

	// La sesión sobrevive fuera del proceso: leemos la fila directo de la base
	var total int
	if err := db.QueryRowContext(ctx,
		`SELECT count(*) FROM sessions WHERE expires_at > now()`).Scan(&total); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Sesiones vivas en la tabla sessions: %d\n", total)
}
