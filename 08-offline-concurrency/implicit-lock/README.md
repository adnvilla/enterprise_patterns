# Implicit Lock

La adquisición y verificación del lock ocurre automáticamente dentro del repositorio (que actúa como unit of work de la sesión): `FindByID` registra la versión leída y `Save` la verifica siempre, así que ningún desarrollador puede olvidar el lock — el `main` nunca menciona versiones ni locks. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Coarse-Grained Lock e Implicit Lock](https://adrianvillafana.com/?p=967).

## Cómo correrlo

    docker compose up -d --wait   # solo si tiene base de datos
    go run .
    docker compose down -v        # limpia el ambiente efímero
