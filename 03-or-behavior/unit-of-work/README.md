# Unit of Work

Mantiene la lista de objetos afectados por una transacción de negocio (nuevos, sucios y eliminados) y escribe todos los cambios juntos en una sola transacción al hacer `Commit`. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Unit of Work](https://adrianvillafana.com/?p=953).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
