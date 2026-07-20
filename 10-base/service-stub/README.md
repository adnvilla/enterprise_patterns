# Service Stub

Un `StubPaymentGateway` en memoria sustituye a la pasarela de pagos real detrás de la misma interfaz: el checkout ni se entera, y el stub además te da el control del guion — puedes provocar el pago rechazado cuando tú decidas (`DeclineOver`) y afirmar sobre lo cobrado (`Charges`). > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Service Stub y Record Set](https://adrianvillafana.com/?p=989).

## Cómo correrlo

    go run .
    go test ./...   # el test del checkout usa el stub
