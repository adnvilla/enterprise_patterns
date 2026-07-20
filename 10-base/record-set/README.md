# Record Set

Filas y columnas genéricas en memoria, sin structs del dominio: `ToRecordSet` lee cualquier `*sql.Rows` hacia `[]map[string]any`, como hacían el RecordSet de ADO y el DataSet de ADO.NET. Para correr sin base de datos, un driver mínimo de `database/sql` sirve filas simuladas en memoria — el `*sql.Rows` que consume el patrón es real. > Parte de la serie [«Patrones de Arquitectura de Aplicaciones Empresariales»](https://adrianvillafana.com/?p=951) — entrada: [Service Stub y Record Set](https://adrianvillafana.com/?p=989).

## Cómo correrlo

    go run .
