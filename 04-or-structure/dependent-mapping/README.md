# Dependent Mapping

El dependiente (`OrderLine`) sí tiene tabla pero no tiene mapper propio: su persistencia la maneja el mapper de su dueño (`OrderMapper`), que lee y escribe las líneas —con borrar e insertar dentro de una transacción— cada vez que lee y escribe el pedido. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Objetos sin tabla propia: Dependent Mapping, Embedded Value y Serialized LOB](https://adrianvillafana.com/?p=987).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
