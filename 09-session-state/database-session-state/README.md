# Database Session State

El estado de sesión se persiste en Postgres como una fila más — tabla `sessions(id, data JSONB, expires_at)` — así que cualquier servidor puede atender cualquier petición y un reinicio no pierde nada. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [¿Dónde vive el estado de sesión?](https://adrianvillafana.com/?p=970).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero

La conexión se toma de `DATABASE_URL` (default `postgres://app:app@localhost:5432/tienda?sslmode=disable`). El main levanta un servidor `net/http` en `:8080`, imprime las rutas de ejemplo para `curl` y hace peticiones de demostración end-to-end contra el Postgres del compose:

    curl -c cookies.txt 'http://localhost:8080/login?customer=7'
    curl -b cookies.txt 'http://localhost:8080/cart/add?product=5'
    curl -b cookies.txt 'http://localhost:8080/cart'
