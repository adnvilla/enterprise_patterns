# Query Object

Representa una consulta como un objeto en términos del dominio —condiciones, orden, límite— que luego un intérprete traduce a SQL con placeholders seguros, ideal para filtros dinámicos sin concatenar strings. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Query Object](https://adrianvillafana.com/?p=966).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
