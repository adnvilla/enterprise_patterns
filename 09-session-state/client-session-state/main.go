package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/cookiejar"

	"github.com/adnvilla/enterprise_patterns/09-session-state/client-session-state/sesion"
)

func main() {
	// El servidor es stateless: todo el estado viaja en la cookie firmada
	http.HandleFunc("/cart/set", func(w http.ResponseWriter, r *http.Request) {
		cartID := r.URL.Query().Get("id")
		if cartID == "" {
			http.Error(w, "falta el parámetro id", http.StatusBadRequest)
			return
		}
		sesion.SetCart(w, cartID)
		fmt.Fprintf(w, "carrito %s guardado en la cookie (firmado)\n", cartID)
	})
	http.HandleFunc("/cart", func(w http.ResponseWriter, r *http.Request) {
		cartID, err := sesion.Cart(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		fmt.Fprintf(w, "tu carrito es %s\n", cartID)
	})

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(ln, nil)

	fmt.Println("Servidor escuchando en :8080")
	fmt.Println("Rutas de ejemplo:")
	fmt.Println(`  curl -c cookies.txt 'http://localhost:8080/cart/set?id=carrito-7'`)
	fmt.Println(`  curl -b cookies.txt 'http://localhost:8080/cart'`)
	fmt.Println(`  curl -H 'Cookie: cart=carrito-999.firma-falsa' 'http://localhost:8080/cart'   # cookie alterada`)
	fmt.Println()

	// Demostración: un cliente con cookie jar
	jar, _ := cookiejar.New(nil)
	client := &http.Client{Jar: jar}

	resp, err := client.Get("http://localhost:8080/cart/set?id=carrito-7")
	if err != nil {
		log.Fatal(err)
	}
	resp.Body.Close()

	resp, err = client.Get("http://localhost:8080/cart")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET /cart con cookie válida  -> %s\n", resp.Status)
	resp.Body.Close()

	// Un cliente tramposo altera el valor: la firma ya no coincide
	req, _ := http.NewRequest("GET", "http://localhost:8080/cart", nil)
	req.Header.Set("Cookie", "cart=carrito-999.firma-falsa")
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("GET /cart con cookie alterada -> %s\n", resp.Status)
	resp.Body.Close()
}
