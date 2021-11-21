# Contenedor de Docker

## Propósito

Crear un entorno de prueba desplegable y repetible de nuestra aplicación para agilizar y facilitar la ejecución de tests.

## Dockerfile

### Base

Se ha utilizado como base la imagen oficial de Golang basada en alpine. Necesitábamos una base que no solo tuviera Go instalado, sino que también tuviera configurado correctamente el entorno de Go. Además, como es lógico, querremos que la imagen sea del menor tamaño posible. Con estas condiciones, surgieron dos candidatos:

- **Imagen oficial de Golang** (debian): Parecía cumplir todos los requisitos pero, la versión por defecto basada en debian llegaba a pesar ~900MB.
- **Imagen oficial de Golang** (alpine): Golang también ofrece una variante de su imagen mucho más ligera (300MB) basada en la distribución Alpine Linux. El único "inconveniente" es que no trae instalado por defecto el paquete de g++, necesario para compilar el proyecto previamente a la ejecución de los tests.

### Dependencias

Será necesario instalar las siguientes dependencias desde el Dockerfile:

- **g++** para el compilador y librerías utilizadas por Go internamente.
- **Task**, el task manager usado en el proyecto.
- **Dependencias del código**: reflejadas en go.mod .

### Buenas prácticas

Se han tenido en cuenta las siguientes buenas prácticas para la realización del Dockerfile:

- **Cambio a un usuario de privilegios reducidos**: Creamos un grupo y un usuario "goautoeq" al que cambiarnos tan pronto como sea posible para no trabajar con el usuario root. 
- **Instalación del menor número de paquetes**: Se han instalado exclusivamente los paquetes necesarios para la realización de la tarea del contenedor.
- **Minimización del número de capas**: Se han realizado pruebas comparando el tamaño final de la imagen agrupando las órdenes RUN y de forma separada. El resultado ha sido que el tamaño final no ha variado nada, por lo que no merece la pena agruparlas. También se han tenido en cuenta las builds multi-stage, pero dada la baja complejidad de la tarea a realizar, no es necesario hacer uso de esto.