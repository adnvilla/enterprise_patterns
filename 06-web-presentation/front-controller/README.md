# Front Controller

Un controlador único recibe todas las peticiones y las despacha a commands según la petición; lo transversal (logging, auth) vive en middlewares escritos una sola vez. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Front Controller](https://adrianvillafana.com/?p=976).

## Cómo correrlo

    go run .

El servidor escucha en `:8080` (sin base de datos, datos en memoria). Prueba con:

    curl -H "Authorization: Bearer demo" "http://localhost:8080/?cmd=orders.list"
    curl -H "Authorization: Bearer demo" "http://localhost:8080/?cmd=orders.show&id=1"
    curl "http://localhost:8080/?cmd=orders.list"   # sin auth → 401
