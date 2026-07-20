# Embedded Value

Aplana un objeto pequeño de valor (`Address`) en columnas de la tabla de su dueño, sin tabla ni fila propia: el objeto existe solo en memoria y el mapper aplana y des-aplana con un `Scan` multi-columna. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Objetos sin tabla propia: Dependent Mapping, Embedded Value y Serialized LOB](https://adrianvillafana.com/?p=987).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
