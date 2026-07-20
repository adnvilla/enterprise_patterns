# Page Controller

Un objeto controlador por cada página o acción del sitio: cada handler interpreta la entrada de SU página, invoca al modelo y elige la vista. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Page Controller](https://adrianvillafana.com/?p=974).

## Cómo correrlo

    go run .

El servidor escucha en `:8080` (sin base de datos, datos en memoria). Prueba con:

    curl http://localhost:8080/orders
    curl "http://localhost:8080/orders/detail?id=1"
