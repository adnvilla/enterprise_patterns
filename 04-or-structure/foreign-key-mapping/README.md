# Foreign Key Mapping

Mapea una referencia entre objetos a una columna de llave foránea: en memoria el `Customer` tiene la colección de `Orders`, y en la base cada fila de `orders` apunta a su `customer_id`; el mapper traduce en ambas direcciones. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Mapeando relaciones: Foreign Key Mapping y Association Table Mapping](https://adrianvillafana.com/?p=983).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
