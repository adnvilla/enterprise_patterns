# Identity Map

Asegura que cada objeto se cargue UNA sola vez por sesión, manteniendo un mapa de `id → objeto`: pedir dos veces el mismo pedido devuelve exactamente la misma instancia en memoria. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Identity Map](https://adrianvillafana.com/?p=957).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
