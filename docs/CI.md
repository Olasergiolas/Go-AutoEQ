# Integración continua

## Github Workflows

- **Publicar imagen Docker a Docker Hub**: Esto se ha conseguido siguiendo la guía de la [documentación oficial de GitHub](https://docs.github.com/en/actions/publishing-packages/publishing-docker-images) con la única diferencia de la adición de los pasos necesarios para hacer uso de una build cache que acelere el proceso de creación de la imagen. Para esto último se ha seguido el apartado de "Optimizing the workflow" de la guía de la [documentación oficial de Docker](https://docs.docker.com/ci-cd/github-actions/) para la creación de un workflow que automatice dicha publicación. El workflow se pondrá en marcha solo cuando se realice un push a la rama main para reducir drásticamente el número de veces que este se lanzaría de hacerse también al realizar un push en las ramas de trabajo para los objetivos.

## CircleCI

- **Automatización de la ejecución de tests**: Para esta tarea se ha elegido utilizar CircleCI debido a su simplicidad para poner en marcha un workflow simple, su soporte out-of-the-box para el uso de imágenes Docker desde DockerHub y su adecuado plan gratuito sin necesidad de introducir un método de pago al contrario que sucede con otro sistemas de CI como Travis. También es necesario comentar que CircleCI presenta algunas desventajas como una documentación que en ocasiones puede ser algo confusa o el tener una comunidad mucho más joven en comparación con las de otras soluciones como Jenkins o Travis.

