# Frameworks de tests y bibliotecas de aserciones

## Framework de tests

### Elección

Este proyecto usará [*Testify*](https://github.com/stretchr/testify) como su framework de tests.

### Motivación

Aunque es cierto que Golang ya posee su propio framework de tests con el paquete *"testing"* he considerdo que podría ser interesante usar Testify en su lugar debido a la facilidad y rapidez que aporta para crear un gran catálogo de aserciones, soporte para mocking y la posibilidad de crear testing suites que permiten realizar el setup y teardown de tests más complejos de una forma más similar a como se hace en algunos OOPL como Dart, donde ya tengo experiencia previa haciendo tests.

## Biblioteca de aserciones

### Elección

La biblioteca de aserciones de Testify.

### Motivación

Debido a que Testify posee su propia biblioteca de aserciones integrada en el framework de tests, no será necesario recurrir a otra externa.

