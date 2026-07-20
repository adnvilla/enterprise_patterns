# Client Session State

El estado de sesión viaja con el cliente: el id del carrito va en una cookie firmada con HMAC, así que el servidor no recuerda nada — solo verifica la firma y lee (firmado no es cifrado: el cliente puede leerla, pero no alterarla). > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [¿Dónde vive el estado de sesión?](https://adrianvillafana.com/?p=970).

## Cómo correrlo

    go run .

El main levanta un servidor `net/http` en `:8080`, imprime las rutas de ejemplo para `curl` y hace peticiones de demostración (incluida una con cookie alterada, que el servidor rechaza):

    curl -c cookies.txt 'http://localhost:8080/cart/set?id=carrito-7'
    curl -b cookies.txt 'http://localhost:8080/cart'
    curl -H 'Cookie: cart=carrito-999.firma-falsa' 'http://localhost:8080/cart'   # cookie alterada
