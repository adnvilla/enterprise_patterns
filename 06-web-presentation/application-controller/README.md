# Application Controller

Centraliza el flujo de navegación en un solo punto: una máquina de estados decide qué pantalla sigue en el checkout (carrito → dirección → pago → confirmación) y los handlers solo delegan. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Application Controller](https://adrianvillafana.com/?p=982).

## Cómo correrlo

    go run .

El servidor escucha en `:8080` (sin base de datos, datos en memoria). Prueba con:

    curl "http://localhost:8080/checkout?state=cart&event=next"
    curl "http://localhost:8080/checkout?state=payment&event=back"
    curl "http://localhost:8080/checkout?state=cart&event=back"   # transición inválida → 400
