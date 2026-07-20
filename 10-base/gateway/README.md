# Gateway

Un `PaymentGateway` encapsula la API HTTP de un proveedor de pagos externo detrás de una interfaz con el vocabulario de la tienda: URLs, headers y formato del proveedor viven en un solo paquete y no se exportan. El main levanta una API externa simulada (`httptest`) para demostrar el cobro aprobado y el rechazado. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Gateway y Mapper](https://adrianvillafana.com/?p=973).

## Cómo correrlo

    go run .
