# enterprise_patterns

Ejemplos en **Go** de los patrones de *Patterns of Enterprise Application Architecture* (PoEAA) de Martin Fowler, acompañando la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) del blog [El Camino del Desarrollador](https://adrianvillafana.com).

Todos los ejemplos comparten el mismo dominio: una pequeña **tienda** (`customers`, `products`, `orders`, `order_lines`).

## Cómo correr un ejemplo

Cada carpeta es un **módulo Go independiente**. Los que usan base de datos incluyen un `docker-compose.yml` con Postgres efímero y su `schema.sql`:

```bash
cd 03-or-behavior/unit-of-work
docker compose up -d --wait   # solo si el ejemplo tiene base de datos
go run .
docker compose down -v        # limpia el ambiente efímero
```

Los ejemplos sin base de datos solo necesitan `go run .`. Los de presentación web levantan un servidor en `:8080` e imprimen las rutas para probar con `curl`.

## Catálogo

### 01 — Lógica de dominio

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [transaction-script](01-domain-logic/transaction-script) | [Transaction Script](https://adrianvillafana.com/?p=958) | ✔ |
| [domain-model](01-domain-logic/domain-model) | [Domain Model](https://adrianvillafana.com/?p=968) | — |
| [table-module](01-domain-logic/table-module) | [Table Module](https://adrianvillafana.com/?p=969) | ✔ |
| [service-layer](01-domain-logic/service-layer) | [Service Layer](https://adrianvillafana.com/?p=972) | ✔ |

### 02 — Fuentes de datos

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [table-data-gateway](02-data-source/table-data-gateway) | [Table Data Gateway y Row Data Gateway](https://adrianvillafana.com/?p=960) | ✔ |
| [row-data-gateway](02-data-source/row-data-gateway) | [Table Data Gateway y Row Data Gateway](https://adrianvillafana.com/?p=960) | ✔ |
| [active-record](02-data-source/active-record) | [Active Record](https://adrianvillafana.com/?p=975) | ✔ |
| [data-mapper](02-data-source/data-mapper) | [Data Mapper](https://adrianvillafana.com/?p=977) | ✔ |

### 03 — Comportamiento objeto-relacional

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [unit-of-work](03-or-behavior/unit-of-work) | [Unit of Work](https://adrianvillafana.com/?p=953) | ✔ |
| [identity-map](03-or-behavior/identity-map) | [Identity Map](https://adrianvillafana.com/?p=957) | ✔ |
| [lazy-load](03-or-behavior/lazy-load) | [Lazy Load](https://adrianvillafana.com/?p=964) | ✔ |

### 04 — Estructura objeto-relacional

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [identity-field](04-or-structure/identity-field) | [Identity Field](https://adrianvillafana.com/?p=980) | ✔ |
| [foreign-key-mapping](04-or-structure/foreign-key-mapping) | [Mapeando relaciones](https://adrianvillafana.com/?p=983) | ✔ |
| [association-table-mapping](04-or-structure/association-table-mapping) | [Mapeando relaciones](https://adrianvillafana.com/?p=983) | ✔ |
| [dependent-mapping](04-or-structure/dependent-mapping) | [Objetos sin tabla propia](https://adrianvillafana.com/?p=987) | ✔ |
| [embedded-value](04-or-structure/embedded-value) | [Objetos sin tabla propia](https://adrianvillafana.com/?p=987) | ✔ |
| [serialized-lob](04-or-structure/serialized-lob) | [Objetos sin tabla propia](https://adrianvillafana.com/?p=987) | ✔ |
| [inheritance-mapping](04-or-structure/inheritance-mapping) | [Mapeando herencia](https://adrianvillafana.com/?p=990) | ✔ |

### 05 — Metadatos y consultas

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [metadata-mapping](05-metadata-queries/metadata-mapping) | [Metadata Mapping](https://adrianvillafana.com/?p=961) | ✔ |
| [query-object](05-metadata-queries/query-object) | [Query Object](https://adrianvillafana.com/?p=966) | ✔ |
| [repository](05-metadata-queries/repository) | [Repository](https://adrianvillafana.com/?p=971) | ✔ |

### 06 — Presentación web

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [mvc](06-web-presentation/mvc) | [Model View Controller](https://adrianvillafana.com/?p=962) | — |
| [page-controller](06-web-presentation/page-controller) | [Page Controller](https://adrianvillafana.com/?p=974) | — |
| [front-controller](06-web-presentation/front-controller) | [Front Controller](https://adrianvillafana.com/?p=976) | — |
| [template-view](06-web-presentation/template-view) | [Template View](https://adrianvillafana.com/?p=959) | — |
| [transform-view](06-web-presentation/transform-view) | [Transform View y Two Step View](https://adrianvillafana.com/?p=979) | — |
| [two-step-view](06-web-presentation/two-step-view) | [Transform View y Two Step View](https://adrianvillafana.com/?p=979) | — |
| [application-controller](06-web-presentation/application-controller) | [Application Controller](https://adrianvillafana.com/?p=982) | — |

### 07 — Distribución

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [remote-facade](07-distribution/remote-facade) | [Remote Facade](https://adrianvillafana.com/?p=965) | — |
| [data-transfer-object](07-distribution/data-transfer-object) | [Data Transfer Object](https://adrianvillafana.com/?p=985) | — |

### 08 — Concurrencia offline

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [optimistic-offline-lock](08-offline-concurrency/optimistic-offline-lock) | [Optimistic Offline Lock](https://adrianvillafana.com/?p=988) | ✔ |
| [pessimistic-offline-lock](08-offline-concurrency/pessimistic-offline-lock) | [Pessimistic Offline Lock](https://adrianvillafana.com/?p=992) | ✔ |
| [coarse-grained-lock](08-offline-concurrency/coarse-grained-lock) | [Coarse-Grained Lock e Implicit Lock](https://adrianvillafana.com/?p=967) | ✔ |
| [implicit-lock](08-offline-concurrency/implicit-lock) | [Coarse-Grained Lock e Implicit Lock](https://adrianvillafana.com/?p=967) | ✔ |

### 09 — Estado de sesión

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [client-session-state](09-session-state/client-session-state) | [¿Dónde vive el estado de sesión?](https://adrianvillafana.com/?p=970) | — |
| [server-session-state](09-session-state/server-session-state) | [¿Dónde vive el estado de sesión?](https://adrianvillafana.com/?p=970) | — |
| [database-session-state](09-session-state/database-session-state) | [¿Dónde vive el estado de sesión?](https://adrianvillafana.com/?p=970) | ✔ |

### 10 — Patrones base

| Ejemplo | Entrada del blog | BD |
|---|---|---|
| [gateway](10-base/gateway) | [Gateway y Mapper](https://adrianvillafana.com/?p=973) | — |
| [mapper](10-base/mapper) | [Gateway y Mapper](https://adrianvillafana.com/?p=973) | — |
| [layer-supertype](10-base/layer-supertype) | [Layer Supertype y Separated Interface](https://adrianvillafana.com/?p=978) | — |
| [separated-interface](10-base/separated-interface) | [Layer Supertype y Separated Interface](https://adrianvillafana.com/?p=978) | — |
| [registry](10-base/registry) | [Registry y Plugin](https://adrianvillafana.com/?p=981) | — |
| [plugin](10-base/plugin) | [Registry y Plugin](https://adrianvillafana.com/?p=981) | — |
| [value-object](10-base/value-object) | [Value Object y Money](https://adrianvillafana.com/?p=984) | — |
| [money](10-base/money) | [Value Object y Money](https://adrianvillafana.com/?p=984) | — |
| [special-case](10-base/special-case) | [Special Case](https://adrianvillafana.com/?p=986) | — |
| [service-stub](10-base/service-stub) | [Service Stub y Record Set](https://adrianvillafana.com/?p=989) | — |
| [record-set](10-base/record-set) | [Service Stub y Record Set](https://adrianvillafana.com/?p=989) | — |

## Serie hermana

Los patrones del Gang of Four viven en [DesignPatterns](https://github.com/adnvilla/DesignPatterns), acompañando la serie [«Patrones de Diseño»](https://adrianvillafana.com/2018/02/19/patrones-de-diseno/).
