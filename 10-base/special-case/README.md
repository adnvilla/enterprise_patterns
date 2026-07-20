# Special Case

Ante un cliente que no existe, el repositorio no devuelve `nil` ni un error: devuelve un `UnknownCustomer` que cumple la misma interfaz con valores seguros («invitado», 0% de descuento) — la decisión de qué significa «no encontrado» se toma una vez, en el proveedor, y el resto del código no repite ifs. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Special Case](https://adrianvillafana.com/?p=986).

## Cómo correrlo

    go run .
