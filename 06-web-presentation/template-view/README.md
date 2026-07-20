# Template View

Renderiza HTML incrustando marcadores en una plantilla estática: `html/template` sustituye cada marcador con datos del view model, y los helpers de formato viven en Go, donde se pueden probar. Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Template View](https://adrianvillafana.com/?p=959).

## Cómo correrlo

    go run .

El servidor escucha en `:8080` (sin base de datos, datos en memoria). La plantilla vive en `vistas/templates/orders.tmpl` y se empaca en el binario con `embed`. Prueba con:

    curl http://localhost:8080/ordenes
