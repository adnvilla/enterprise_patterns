# Data Transfer Object

Structs planos con tags `json` que definen el contrato público de la API, más un assembler que traduce entre el DTO y el dominio: los campos internos (`Version`, `CostCents`) nunca salen y los protegidos (`status`) nunca entran. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Data Transfer Object](https://adrianvillafana.com/?p=985).

## Cómo correrlo

    go run .

El main levanta un servidor `net/http` en `:8080`, imprime la ruta de ejemplo para `curl` y hace una petición de demostración:

    curl -X POST http://localhost:8080/orders -H 'Content-Type: application/json' \
      -d '{"customer_id":7,"lines":[{"product_id":1,"quantity":2}]}'
