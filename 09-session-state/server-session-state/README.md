# Server Session State

El estado de sesión vive en la memoria del servidor (un mapa protegido con `sync.RWMutex` y con TTL); el cliente solo lleva un identificador opaco en la cookie. Es el mismo patrón que Redis o memcached, en su versión más simple. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [¿Dónde vive el estado de sesión?](https://adrianvillafana.com/?p=970).

## Cómo correrlo

    go run .

El main levanta un servidor `net/http` en `:8080`, imprime las rutas de ejemplo para `curl` y hace peticiones de demostración (incluida una con un id de sesión desconocido, que el servidor rechaza):

    curl -c cookies.txt 'http://localhost:8080/login?customer=7'
    curl -b cookies.txt 'http://localhost:8080/cart/add?product=5'
    curl -b cookies.txt 'http://localhost:8080/cart'
