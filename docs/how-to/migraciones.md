# Migraciones

Esta sección contiene información sobre cómo realizar migraciones en el proyecto.

Una migración es un cambio en la estructura de la base de datos que se aplica en orden secuencial, estas se encuentran en la carpeta `internal/infrastructure/postgres/migrations`.

## Crear un nuevo archivo de migración

Generar un nuevo archivo de migración de forma secuencial en la carpeta usada:

```sh
migrate create \
    -seq \
    -ext sql \
    --dir internal/infrastructure/postgres/migrations \
    migrationidentifier
```

## Aplicar migraciones

Para aplicar migraciones, se puede usar el comando `migrate` de la siguiente manera:

```sh
migrate \
    -database "postgres://postgres:postgres@127.0.0.1:5432/99minutos?sslmode=disable" \
    -path "internal/infrastructure/postgres/migrations" \
    up
```

## Deshacer una migración

Para deshacer migraciones, se puede usar el comando `migrate` de la siguiente manera:

```sh
migrate \
    -database "postgres://postgres:postgres@127.0.0.1:5432/99minutos?sslmode=disable" \
    -path "internal/infrastructure/postgres/migrations" \
    down 1
```
