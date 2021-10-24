# <img src="https://pbs.twimg.com/profile_images/378800000364886592/2e6f7c62714a4ae59c71e7cc8772df0e_400x400.png" alt="logo" width=50px /> Go-AutoEQ

Repositorio para almacenar el proyecto a realizar para la asignatura de Infraestructura Virtual en el curso 21-22.

## Tabla de contenidos :card_index_dividers:

- [Idea](#idea-bulb)
- [Instalación](#instalación-zap)
- [Gestor de tareas](#gestor-de-tareas-rocket)
- [Documentación](#documentación-memo)

## Idea :bulb:

Facilitar el proceso de obtener perfiles de ecualización recomendados para auriculares que permitan mejorar la experiencia de sonido del usuario en Linux. Para esto se utilizaría la información recopilada en el repositorio de [AutoEQ](https://github.com/jaakkopasanen/AutoEq), que contiene los hallazgos de diversos investigadores de la materia.

## Instalación :zap:

#### 1. Pasos previos

Antes de comenzar la instalación, es necesario tener Go previamente instalado en tu sistema. Para más información sobre cómo hacerlo accede [aquí](https://golang.org/doc/install).

#### 2. Instalación de task

Para instalar el gestor de tareas ***task*** en el directorio /usr/local/bin:

```shell
sh -c "$(curl --location https://taskfile.dev/install.sh)" -- -d -b /usr/local/bin
```

#### 3. Clonado del repositorio

Para clonar este repositorio en su sistema ejecute la siguiente orden:

```shell
git clone https://github.com/Olasergiolas/Go-AutoEQ.git
```

Ahora puedes acceder al nuevo directorio creado con:

```shell
cd Go-AutoEQ/
```

#### 4. Instalación del proyecto

Por último, para instalar Go-AutoEQ solo es necesario ejecutar:

```shell
task install
```

Nota: Go-AutoEQ será instalado en `go env GOPATH`/bin y es posible que este directorio no esté incluido en el $PATH de tu sistema.

## Gestor de tareas :rocket:

Además del `task install` también están disponibles una serie de tareas como:

- `task test`: Pondrá en marcha los tests del proyecto.
- `task check`: Comprobará si existe algún error de sintaxis intentando compilar el proyecto.
- `task installdeps`: Actualizará las dependencias del proyecto en el fichero *go.mod* y los checksum de estas en el fichero *go.sum* . 
- `task fmt`: Dará formato a todos los archivos de código .go del proyecto para hacerlos más fácil de leer y mantener.

## Documentación :memo:

Puede acceder a la documentación del proyecto desde [aquí](docs).

