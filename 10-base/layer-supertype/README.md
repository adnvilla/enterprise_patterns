# Layer Supertype

Un struct `Entity` embebido concentra el comportamiento común de toda la capa de dominio (el manejo del ID y `IsNew`): `Order` y `Customer` lo obtienen con una línea de embebido — composición en lugar de herencia. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Layer Supertype y Separated Interface](https://adrianvillafana.com/?p=978).

## Cómo correrlo

    go run .
