# Lazy Load

Un objeto que no contiene todos los datos que necesitas, pero sabe cómo obtenerlos: el pedido se carga sin sus líneas y estas se leen de la base de datos solo en el primer acceso a `Lines()` (lazy initialization con `sync.Once`). Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Lazy Load](https://adrianvillafana.com/?p=964).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
