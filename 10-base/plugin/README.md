# Plugin

El mismo binario se comporta distinto por entorno: la variable `PAYMENT_GATEWAY` (configuración, no compilación) elige al arrancar cuál implementación del `PaymentGateway` se enlaza — la real o la sandbox — a través de un registry consultado una sola vez. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Registry y Plugin](https://adrianvillafana.com/?p=981).

## Cómo correrlo

    go run .                        # usa la implementación sandbox (default)
    PAYMENT_GATEWAY=real go run .   # misma compilación, otra implementación
