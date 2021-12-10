# Integración continua

## Publicar imagen Docker a Docker Hub

Esto se ha conseguido haciendo uso de un **Github Workflow** creado siguiendo la guía de la [documentación oficial de GitHub](https://docs.github.com/en/actions/publishing-packages/publishing-docker-images) con la única diferencia de la adición de los pasos necesarios para hacer uso de una build cache que acelere el proceso de creación de la imagen. Para esto último se ha seguido el apartado de "Optimizing the workflow" de la guía de la [documentación oficial de Docker](https://docs.docker.com/ci-cd/github-actions/) para la creación de un workflow que automatice dicha publicación. El workflow se pondrá en marcha solo cuando se realice un push a la rama main para reducir drásticamente el número de veces que este se lanzaría de hacerse también al realizar un push en las ramas de trabajo para los objetivos.

## Automatización de la ejecución de tests

Para esta tarea necesitaremos elegir dos sistemas de integración continua que han de satisfacer los siguientes requisitos:

- **Integración con Docker Hub**: Posibilidad de hacer uso de imágenes docker desde Docker Hub de forma nativa y sencilla, para facilitar la puesta en marcha del contenedor de pruebas del proyecto.
- **Soporte para matrix jobs**: Con el objetivo de poder ejecutar los tests en las dos versiones de Go con soporte actualmente.
- **Soporte para Github Checks**: Para poder gestionar el workflow desde Github y obtener mensajes de estado de este.
- **Sin necesidad de instalación**: Querremos que no sea necesario tener que hostear, instalar y configurar por nuestra cuenta el sistema de integración continua en un servidor.
- **Gratuito sin necesidad de introducir un método de pago**

Considerando algunos de los sistemas de integración continua más populares como son Travis CI, Jenkins, TeamCity (Jetbrains), Circle CI y Github actions y teniendo en cuenta los requisitos anteriores, elegimos **Circle CI** y **Github actions** ya que son los únicos de los mencionados que cumplen todos los requisitos especificados. En estos sistemas de CI realizamos las siguientes tareas:

- **Github actions**: Poner en marcha los tests comprobando su correcto funcionamiento en las dos "major releases" actuales de Go mediante una job matrix. El ciclo de releases de Go funciona de forma que cada versión "major" tiene soporte hasta que se publican dos nuevas "major releases". También cabe destacar que Go no ofrece una versión LTS, así que siempre nos interesará probar solo las dos versiones con soporte actualmente.
- **Circle CI**: Poner en marcha los tests haciendo uso del contenedor de pruebas Docker publicado en Docker Hub.

