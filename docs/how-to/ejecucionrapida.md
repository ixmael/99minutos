# Ejecución rápida con docker-compose

Para ejecutar el proyecto de forma rápida utilizando `docker-compose`, siga estos pasos:

1. Tener Docker y docker-compose instalados.
2. Navega hasta el directorio raíz del proyecto.
3. Copiar el archivo `.env.toml.example` a `.env.toml`.
4. Ejecuta el siguiente comando para iniciar los contenedores:

```bash
docker-compose up -d
```

4. Una vez que los contenedores estén en ejecución, puedes acceder a la aplicación en tu navegador visitando `http://localhost:8080`.

5. Verificar que la aplicación está funcionando correctamente.

```bash
curl -I http://localhost:8080/health
```
