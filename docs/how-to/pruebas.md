# ¿Cómo ejecutar las pruebas?

En esta sección encontrarás guías paso a paso para trabajar con las pruebas.

## Pruebas de integración

Las pruebas de integración son descritas con **sintaxis Gherkin** en el directorio `test/features`.

Puede ejecutar las pruebas de integración utilizando el siguiente comando:

```sh
go test -tags=integration ./test/bdd
```
