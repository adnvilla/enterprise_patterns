package sesion

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/http"
	"strings"
)

var secret = []byte("clave-secreta-de-la-tienda")

func sign(value string) string {
	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(value))
	return base64.RawURLEncoding.EncodeToString(mac.Sum(nil))
}

// Client Session State: el estado viaja en la cookie, firmado
func SetCart(w http.ResponseWriter, cartID string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "cart",
		Value:    cartID + "." + sign(cartID),
		HttpOnly: true,
	})
}

// El servidor no recuerda nada: solo verifica la firma
func Cart(r *http.Request) (string, error) {
	c, err := r.Cookie("cart")
	if err != nil {
		return "", err
	}
	parts := strings.SplitN(c.Value, ".", 2)
	if len(parts) != 2 || !hmac.Equal([]byte(parts[1]), []byte(sign(parts[0]))) {
		return "", errors.New("cookie alterada")
	}
	return parts[0], nil
}
