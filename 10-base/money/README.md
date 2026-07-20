# Money

El value object insignia del libro: dinero como `int64` en centavos más moneda, con `Add` (que rechaza mezclar monedas), `Multiply` y `Allocate` — el reparto en n partes que nunca pierde un centavo (10 entre 3 da 4, 3 y 3; no 3.33 × 3). > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Value Object y Money](https://adrianvillafana.com/?p=984).

## Cómo correrlo

    go run .
    go test ./...   # incluye el test de Allocate
