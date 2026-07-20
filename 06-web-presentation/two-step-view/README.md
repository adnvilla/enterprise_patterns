# Two Step View

Renderiza en dos pasos: primero una representación lógica de la pantalla (qué se muestra) y después un layout común a todo el sitio que la convierte en HTML (cómo se ve); rediseñar el sitio es tocar solo `templates/layout.tmpl`. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Transform View y Two Step View](https://adrianvillafana.com/?p=979).

## Cómo correrlo

    go run .

El servidor escucha en `:8080` (sin base de datos, datos en memoria). Las plantillas viven en `templates/` (`layout.tmpl` + `orders.tmpl`) y se empacan en el binario con `embed`. Prueba con:

    curl http://localhost:8080/ordenes
