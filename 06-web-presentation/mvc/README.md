# Model View Controller

Divide la interfaz en tres roles: un modelo con los datos y la lógica de dominio, una vista que lo presenta y un controlador que recibe la entrada y coordina a los otros dos. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Model View Controller](https://adrianvillafana.com/?p=962).

## Cómo correrlo

    go run .

El servidor escucha en `:8080` (sin base de datos, datos en memoria). Prueba con:

    curl http://localhost:8080/orders
