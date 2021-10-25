# Gestores de tareas y dependencias

## Gestor de tareas

### Elección

Este proyecto usará [Task](https://github.com/go-task/task) como su task runner.

### Motivación

Aunque Go ofrece su propio task runner implícito mediante subcomandos como "go test", este no permite definir nuevas tareas más allá de las que se ofrecen por defecto, además de que para muchas de estas sería necesario especificar rutas concretas del proyecto para su correcta ejecución. Es por esto que se ha decidido utilizar otro task runner que sí permita especificar nuevas tareas manualmente.

Los factores determinantes para la elección de Task para cumplir esta tarea han sido su simplicidad de uso, su correcta documentación, su sencillo proceso de instalación y el hecho de estar también escrito en Go.

## Gestor de dependencias

### Elección

Se usará el gestor de dependencias propio de [Go](https://golang.org/doc/modules/managing-dependencies).

### Motivación

Este gestor ya integrado no presenta problemas de uso y aporta todo lo necesario para este proyecto.

