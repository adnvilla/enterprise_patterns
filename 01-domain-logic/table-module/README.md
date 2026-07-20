# Table Module

Una sola instancia (`OrdersModule`) concentra la lógica de negocio de TODAS las filas de la tabla `orders`: totales por cliente y marcado en bloque de pedidos vencidos, operando sobre conjuntos con SQL. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Table Module](https://adrianvillafana.com/?p=969).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
