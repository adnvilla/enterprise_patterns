# Repository

Media entre el dominio y la capa de datos con una interfaz estilo colección en memoria, definida desde el dominio: el negocio pide «los pedidos pagados de este cliente» y las implementaciones (Postgres o memoria) resuelven el cómo. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Repository](https://adrianvillafana.com/?p=971).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero

Los tests de la lógica de dominio usan la implementación en memoria y no necesitan base de datos:

    go test ./...
