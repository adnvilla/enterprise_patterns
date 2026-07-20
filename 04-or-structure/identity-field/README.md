# Identity Field

Guarda el id de la fila de la base de datos dentro del objeto en memoria para mantener la correspondencia objeto ↔ fila: la base genera la clave sustituta con `BIGSERIAL` y el mapper la deposita en el struct con `INSERT ... RETURNING`. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Identity Field](https://adrianvillafana.com/?p=980).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
