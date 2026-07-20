# Active Record

Un objeto que envuelve una fila de la base de datos, encapsula el acceso a ella y agrega lógica de dominio: el propio `Order` sabe insertarse, actualizarse, borrarse y marcar el pedido como pagado. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Active Record](https://adrianvillafana.com/?p=975).

## Cómo correrlo

    docker compose up -d --wait   # solo si tiene base de datos
    go run .
    docker compose down -v        # limpia el ambiente efímero
