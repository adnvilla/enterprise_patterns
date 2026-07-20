# Association Table Mapping

Resuelve las relaciones muchos-a-muchos con una tabla intermedia de dos llaves foráneas (`product_tags`) que no tiene objeto en memoria: al leer se hace JOIN a través de ella y al escribir se sincronizan sus filas con borrar e insertar dentro de una transacción. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Mapeando relaciones: Foreign Key Mapping y Association Table Mapping](https://adrianvillafana.com/?p=983).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
