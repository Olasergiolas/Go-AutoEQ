# Registro de eventos :memo:
## Paquete de logging utilizado
Finalmente se ha decidido utilizar [zerolog](https://github.com/rs/zerolog).
## Motivación
En primer lugar, es necesario definir qué requisitos esperamos de los distintos servicios que van a ser comparados. El servicio de logging que utilicemos para este proyecto tendrá que cumplir lo siguiente:
- **Posibilidad de usar formato JSON**: Esto resulta muy conveniente para poder parsear los logs en otras herramientas.
- **Posibilidad de establecer campos por defecto**: Querremos que los logs contengan campos por defecto como el timestamp o la versión de Go utilizada para no tener que indicar de forma repetitiva que se incluyan dichos campos.
- **Soporte para cualquier io.Writer**: Querremos poder escribir los logs tanto en la salida estandar (os.Stdout) como en un fichero.
- **Distintos niveles de log**: Posibilidad de establecer el nivel de logging deseado (Info, Warn, Error, Fatal...).

Respecto su estado de mantenimiento, requerimos lo básico en su repositorio:
- Que el proyecto haya tenido commits en el último año.
- Que no tenga un número considerable de PRs o Issues acumulados sin atender.
- Que el desarrollo del proyecto no esté detenido.

Teniendo todos estos requisitos en cuenta, comparamos **log**, **logrus**, **zerolog** y **apex log**.
### Log
Es el paquete de logging ya incluido por defecto en Go. Queda descartado debido a que carece de soporte para logs con formato o niveles de logging.
### Logrus
Es el más comúnmente utilizado. Aunque cumple los requisitos de funcionalidad, su desarrollo ha quedado reducido a arreglar pequeños problemas ya que son incapaces de realizar cambios considerables sin romper la compatibilidad con proyectos que lo utilizan. Debido a esto, ya que es factible que una versión futura de Go deje de ser compatible con logrus, descartamos esta opción.
### Apex log
Es también una opción popular para implementar el uso de logs en Go, pero por desgracia, su repositorio de Github parece algo abandonado, con issues y PRs abiertas desde hace meses sin actividad, además de no haber tenido commits desde hace más de un año.

### Zerolog
Zerolog se caracteriza por su velocidad y la posibilida de generar logs con formato. Zerolog cumple con todos los requisitos expuestos previamente, así que será nuestra elección.

## Implementación
Para la implementación, se han tenido en cuenta algunas de las mejores prácticas de logging como la implementación de una interfaz de logging estándar que haga uso de mensajes y eventos predefinidos, inyección de dependencias, escritura a archivo o el uso de logs formateados.