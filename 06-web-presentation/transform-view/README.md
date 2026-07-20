# Transform View

Produce la salida transformando el modelo elemento por elemento: cada `Order` del dominio pasa por un DTO que define la forma pública de la respuesta JSON. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Transform View y Two Step View](https://adrianvillafana.com/?p=979).

## Cómo correrlo

    go run .

El servidor escucha en `:8080` (sin base de datos, datos en memoria). Prueba con:

    curl http://localhost:8080/api/ordenes
