# Remote Facade

Una fachada gruesa (*coarse-grained*) sobre un modelo de dominio fino: el cliente coloca un pedido completo —cliente, líneas y pago— en UNA sola petición HTTP, y por dentro la fachada coordina el modelo con llamadas locales baratas. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Remote Facade](https://adrianvillafana.com/?p=965).

## Cómo correrlo

    go run .

El main levanta un servidor `net/http` en `:8080`, imprime la ruta de ejemplo para `curl` y hace una petición de demostración:

    curl -X POST http://localhost:8080/orders/place -H 'Content-Type: application/json' \
      -d '{"customer_id":7,"lines":[{"product_id":1,"quantity":2,"unit_cents":12900}],"payment_cents":25800}'
