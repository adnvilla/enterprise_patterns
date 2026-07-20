# Metadata Mapping

Guarda el mapeo objeto-relacional —qué tabla, qué columna corresponde a qué campo— como metadatos (un mapa explícito o struct tags) y genera el SQL con una sola pieza de código genérica, en lugar de repetirlo en cada mapper. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Metadata Mapping](https://adrianvillafana.com/?p=961).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
