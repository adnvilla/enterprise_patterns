package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"

	"github.com/adnvilla/enterprise_patterns/09-session-state/server-session-state/sesion"
)

const ttl = 30 * time.Minute

// El cliente solo lleva un identificador opaco
func newSessionID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func main() {
	store := sesion.NewStore()

	// Crea la sesión en el servidor y entrega solo el id en la cookie
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		customerID, _ := strconv.ParseInt(r.URL.Query().Get("customer"), 10, 64)
		id := newSessionID()
		store.Put(id, sesion.Session{CustomerID: customerID}, ttl)
		http.SetCookie(w, &http.Cookie{Name: "sid", Value: id, HttpOnly: true})
		fmt.Fprintf(w, "sesión creada para el cliente %d\n", customerID)
	})

	// Modifica el estado que vive en el servidor
	http.HandleFunc("/cart/add", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sid")
		if err != nil {
			http.Error(w, "sin sesión", http.StatusUnauthorized)
			return
		}
		sess, ok := store.Get(c.Value)
		if !ok {
			http.Error(w, "sesión expirada o desconocida", http.StatusUnauthorized)
			return
		}
		productID, _ := strconv.ParseInt(r.URL.Query().Get("product"), 10, 64)
		sess.CartItems = append(sess.CartItems, productID)
		store.Put(c.Value, sess, ttl)
		fmt.Fprintf(w, "producto %d agregado; carrito: %v\n", productID, sess.CartItems)
	})

	// Lee el estado que vive en el servidor
	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("sid")
		if err != nil {
			http.Error(w, "sin sesión", http.StatusUnauthorized)
			return
		}
		sess, ok := store.Get(c.Value)
		if !ok {
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

	// Demostración: el cliente solo carga la cookie con el id opaco;
	// los datos reales viven en el mapa del servidor (con TTL)
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

	// Con un id desconocido no hay nada que leer: la sesión vive en el servidor
	req, _ := http.NewRequest("GET", "http://localhost:8080/cart", nil)
	req.Header.Set("Cookie", "sid=id-inventado")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET /cart con sid desconocido                    -> %s\n", resp.Status)
	resp.Body.Close()
}
