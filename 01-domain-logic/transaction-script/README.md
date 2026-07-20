# Transaction Script

Organiza la lógica de dominio como un procedimiento por cada transacción de negocio: la función `PlaceOrder` valida, calcula el total y persiste el pedido de principio a fin, dentro de una sola transacción. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Transaction Script](https://adrianvillafana.com/?p=958).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
