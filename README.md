# 99minutos test

Este es el proyecto de prueba para el desafío de **99minutos**.

La documentación se basa en [Diataxis Framework](https://diataxis.fr/) y se encuentra en la carpeta `docs`.

## Contenido

- [Guías](docs/how-to/index.md)
  - [¿Cómo ejecutar pruebas?](docs/how-to/pruebas.md)
  - [¿Cómo aplicar migraciones?](docs/how-to/migraciones.md)
  - [¿Cómo configurar el entorno?](docs/how-to/preparacion.md)
  - [Ejecución rápida con docker-compose](docs/how-to/ejecucionrapida.md)
- [Tutoriales](docs/tutorials/index.md)
  - [Registrar un usuario](docs/tutorials/registrousuario.md)
- [Reference](docs/reference/index.md)
- [Explicación](docs/explanation/index.md)

## Resumen

Este resumen está basado en el contenido de la documentación y se basa a lo solicitado en la prueba.

1. [Diagrama de arquitectura del sistema](docs/assets/architecture.png).
2. Decisiones de diseño: ¿por qué elegiste esa base de datos? ¿Cómo estructuraste el código y
   por qué?

- Arquitectura ports/adapters para el código
  - Tener la lógica de negocio centralizada `internal/core/domain`
  - Poder implementar varias aplicaciones como una **TUI**, una **CLI** o una aplicación **GRPC**.
- PostreSQL
  - Mantener consistencia de datos
  - Mantener integridad referencial
- Valkey
  - Alternativo open source a Redis
  - Las pruebas de rendimiento parecen sobrepasar la latencia de Redis
- RabbitMQ
  - Es una aplicación más simple que Kafka

3. Trade-offs: ¿qué simplificaste y cómo lo resolverías en producción?

- El manejo de errores y excepciones:
  - Definir errores de dominio, de aplicación y de infraestructura
- Mejorar logs con:
  - Rastrear todo el proceso con un identificador único
- Manejo simple de los "roles"
  - Implementar un sistema de roles basado en permisos y privilegios
- Guardado de los eventos de los envíos recibidos

4. Escalabilidad: ¿cómo adaptarías este sistema para manejar 10,000 eventos por segundo?

- Separar el worker en microservicios independientes
- Implementar escalamiento horizontal de los microservicios
- Tener un balanceador de carga para los microservicios
- Mantener los microservicios en una red interna
- Tener un API Gateway público para acceder a la Rest API
- Tener un WAF para limpiar datos maliciosos
