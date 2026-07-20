# Registry

Un `Registry` es el objeto bien conocido donde se encuentran los servicios: un struct con un mapa protegido por `sync.RWMutex`, creado en `main` y consultado una sola vez al arrancar — de ahí en adelante, inyección explícita del servicio ya resuelto. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Registry y Plugin](https://adrianvillafana.com/?p=981).

## Cómo correrlo

    go run .
