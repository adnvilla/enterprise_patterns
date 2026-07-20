# Pessimistic Offline Lock

Evita los conflictos por adelantado: antes de editar un pedido hay que adquirir un lock exclusivo en la tabla `locks` (un `INSERT ... ON CONFLICT DO NOTHING` arbitrado por la clave primaria); el `main` muestra al segundo usuario rechazado mientras el primero tiene reservado el pedido. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Pessimistic Offline Lock](https://adrianvillafana.com/?p=992).

## Cómo correrlo

    docker compose up -d --wait   # solo si tiene base de datos
    go run .
    docker compose down -v        # limpia el ambiente efímero
