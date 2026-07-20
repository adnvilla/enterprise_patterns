# Value Object

`Money` es un objeto cuya identidad ES su valor: struct pequeño e inmutable (receptores por valor, operaciones que devuelven un Money nuevo), comparable con `==`, con sus reglas de negocio — como «no mezcles monedas» — escritas una sola vez dentro del tipo. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Value Object y Money](https://adrianvillafana.com/?p=984).

## Cómo correrlo

    go run .
