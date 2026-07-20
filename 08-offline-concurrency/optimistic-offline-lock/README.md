# Optimistic Offline Lock

Previene conflictos entre transacciones de negocio concurrentes detectándolos al confirmar: cada pedido lleva una columna `version` y el `UPDATE ... WHERE version = $n` solo escribe si nadie se adelantó; el `main` simula dos usuarios editando el mismo pedido y muestra el conflicto detectado. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Optimistic Offline Lock](https://adrianvillafana.com/?p=988).

## Cómo correrlo

    docker compose up -d --wait   # solo si tiene base de datos
    go run .
    docker compose down -v        # limpia el ambiente efímero
