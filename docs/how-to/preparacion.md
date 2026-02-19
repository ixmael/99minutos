# ¿Cómo configurar el entorno?

Para preparar la aplicación y ejecutarla, necesita definir un archivo con formato `TOML`. La aplicación por defecto busca el archivo `.env.toml`, pero si desea usar otro archivo, puede especificarlo en la línea de comandos usando la opción `--config` o `-c`.

El archivo de configuración tiene la siguiente estructura:

```toml
environment="develop"

[restapi]
port = 8080

[repository]
postgres = "postgres://postgres:postgres@127.0.0.1:5432/99minutos?sslmode=disable"
```

En el repositorio puede encontrar un archivo base llamado `.env.toml.example`.
