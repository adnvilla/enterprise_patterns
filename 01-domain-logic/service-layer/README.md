# Service Layer

Una fachada de casos de uso (`OrderService`) define el límite transaccional y coordina el Domain Model con la infraestructura; el handler HTTP solo traduce la petición y delega. El `main` levanta el servidor y hace una petición de demostración end-to-end. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Service Layer](https://adrianvillafana.com/?p=972).

## Cómo correrlo

    docker compose up -d --wait
    go run .
    docker compose down -v        # limpia el ambiente efímero
