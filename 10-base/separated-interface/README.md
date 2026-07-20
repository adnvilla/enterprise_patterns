# Separated Interface

El paquete de dominio define la interfaz `OrderRepository` (el consumidor pone el contrato) y los paquetes de infraestructura la implementan: `postgres` con SQL y `memoria` sin base de datos — la flecha de dependencia siempre apunta al dominio, y solo el `main` conecta ambos mundos. El main corre con la implementación en memoria. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Layer Supertype y Separated Interface](https://adrianvillafana.com/?p=978).

## Cómo correrlo

    go run .
