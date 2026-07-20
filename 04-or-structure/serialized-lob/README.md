# Serialized LOB

Serializa un grafo pequeño de objetos (`Specs`) como un solo blob JSON con `encoding/json` y lo guarda en una columna `JSONB` de la tabla dueña: la estructura variable entra y sale completa de una columna, sin migraciones por cada campo nuevo. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Objetos sin tabla propia: Dependent Mapping, Embedded Value y Serialized LOB](https://adrianvillafana.com/?p=987).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
