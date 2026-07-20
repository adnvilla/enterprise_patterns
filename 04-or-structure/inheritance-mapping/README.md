# Inheritance Mapping (Single, Class y Concrete Table)

Un solo ejemplo que demuestra las tres estrategias de mapeo de herencia sobre la MISMA jerarquía (`Product` con `PhysicalProduct` y `DigitalProduct`): *Single Table* (una tabla con discriminador), *Class Table* (tabla por clase unidas por id) y *Concrete Table* (tabla por clase concreta, polimorfismo vía `UNION`), organizadas con *Inheritance Mappers*. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Mapeando herencia](https://adrianvillafana.com/?p=990).

El `schema.sql` trae las tres variantes de tablas y el `main` inserta y lee los mismos dos productos con cada estrategia, comparándolas.

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
