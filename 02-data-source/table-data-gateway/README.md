# Table Data Gateway

Un objeto que actúa como puerta de acceso a una tabla completa: `OrdersGateway` concentra todo el SQL de la tabla `orders` y devuelve datos planos, sin lógica de negocio. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Table Data Gateway y Row Data Gateway](https://adrianvillafana.com/?p=960).

## Cómo correrlo

    docker compose up -d --wait   # solo si tiene base de datos
    go run .
    docker compose down -v        # limpia el ambiente efímero
