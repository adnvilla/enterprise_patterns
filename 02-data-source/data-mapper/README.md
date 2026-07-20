# Data Mapper

Una capa de mappers que mueve datos entre los objetos de dominio y la base de datos, manteniendo a ambos ignorantes el uno del otro: el package `dominio` no importa `database/sql` y todo el SQL vive en `OrderMapper`. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Data Mapper](https://adrianvillafana.com/?p=977).

## Cómo correrlo

    docker compose up -d --wait   # solo si tiene base de datos
    go run .
    docker compose down -v        # limpia el ambiente efímero
