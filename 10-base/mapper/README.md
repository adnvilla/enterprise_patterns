# Mapper

El mapper `toCharge` traduce entre el modelo del proveedor de pagos (`providerResponse`, que no se exporta) y el `Charge` del dominio: ninguno de los dos subsistemas sabe que el otro existe ni que alguien está traduciendo — un tercero invoca la traducción. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Gateway y Mapper](https://adrianvillafana.com/?p=973).

## Cómo correrlo

    go run .
