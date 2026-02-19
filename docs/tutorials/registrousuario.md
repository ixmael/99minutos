# Crear usuario

Para comenzar a usar el sistema necesita crear una cuenta. Este tutorial le guiará a través del proceso de creación de una cuenta.

Puede usar la colección de [colección de Bruno](collections/99minutosTest/bruno.json) o [colección de Postman](collections/99minutosTest.postman.json).

## Crear usuarios

El sistema permite crear dos tipos de usuarios: administradores y clientes.

### Crear un administrador

Hacer una petición al siguiente endpoint: `POST /users` con los datos del usuario en el cuerpo de la petición en formato JSON, ejemplo:

```json
{
  "email": "admin@ixmael.dev",
  "password": "12345678",
  "is_admin": true
}
```

### Crear un cliente

Hacer una petición al siguiente endpoint: `POST /users` con los datos del usuario en el cuerpo de la petición en formato JSON, ejemplo:

```json
{
  "email": "client@ixmael.dev",
  "password": "12345678"
}
```

## Acceso

Para obtener un token de acceso haga una petición al siguiente endpoint: `POST /auth` con los datos del usuario en el cuerpo de la petición en formato JSON, ejemplo:

```json
{
  "email": "client@ixmael.dev"
  "password": "12345678"
}
```

que como resultado le da el siguiente ejemplo:

```json
{
  "token": "eyJhb...HBSI"
}
```

el valor del token la puede usar para autenticarse en las demás peticiones.
