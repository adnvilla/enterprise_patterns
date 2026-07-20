# Coarse-Grained Lock

Un solo lock protege un conjunto de objetos relacionados: la columna `version` vive únicamente en la raíz del agregado (`orders`) y cubre también sus `order_lines`, así que editar cualquier parte del pedido excluye ediciones concurrentes del todo. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Coarse-Grained Lock e Implicit Lock](https://adrianvillafana.com/?p=967).

## Cómo correrlo

    docker compose up -d --wait   # solo si tiene base de datos
    go run .
    docker compose down -v        # limpia el ambiente efímero
