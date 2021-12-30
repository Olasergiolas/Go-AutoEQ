# Configuración remota :gear:
## Sistema elegido
Para esta tarea, se han utilizado [etcd](https://etcd.io/) y [godotenv](https://github.com/joho/godotenv).
## Motivación
### Almacenamiento distribuido
Para implementar la funcionalidad de configurar la aplicación de forma remota, necesitamos algún sistema de almacenamiento distribuido que no solo nos permita poder almacenar ciertas configuraciones, sino que también se encargue automáticamente de comunicar al resto de nodos los posibles cambios de configuración realizados, de forma que sea posible escalar el número de instancias de la aplicación y que todas compartan una misma configuración. Hay una serie de servicios que pueden proveernos con la funcionalidad descrita como son **etcd**, **Zookeeper** y **consul**.

El servicio finalmente elegido deberá de ser fácil de configurar, usar un mecanismo de almacenamiento simple (más cercano al almacenamiento de parejas clave-valor que a bases de datos SQL) y tener un cliente nativo para Go.
#### 1. etcd
Es un sistema de almacenamiento distribuido de parejas clave-valor escrito en Go con la fiabilidad y tolerancia a fallos como principales objetivos. Gracias a su simplicidad, facilidad de configuración y su [cliente nativo](https://pkg.go.dev/go.etcd.io/etcd/clientv3) de Go, [etcd](https://etcd.io/) es una alternativa a considerar.
#### 2. ZooKeeper
[ZooKeeper](https://zookeeper.apache.org/) es el proyecto de coordinación distribuida de Apache usado en algunos de sus proyectos como Kafka o Solr. Para nuestro uso, tiene las desventajas de no tener un cliente oficial para Go (aunque existe uno mantenido por la [comunidad](https://github.com/go-zookeeper/zk) pero cuya documentación carece de ejemplos de uso rápidos que nos faciliten su puesta en marcha), estar escrito en Java, o el que ZooKeeper tenga que abrir un nuevo socket de conexión por cada petición que realicemos.
#### 3. Consul
[Consul](https://www.consul.io/) es la solución diseñada por Hashicorp. Aunque cumple con los requisitos de funcionalidad establecidos (cliente oficial en Go, almacenamiento KV y fácil configuración), es un servicio que aporta mucha más funcionalidad de la que necesitamos para este proyecto. Es por esto que nos inclinamos por un servicio más simple y orientado a nuestras necesidades como es **etcd**.

### Carga de variables de entorno
Además, por motivos de seguridad y de buenas prácticas, querremos poder tener una forma de evitar hardcodear en el cliente parámetros como por ejemplo la dirección IP o claves de acceso al servidor etcd. Esto se puede conseguir haciendo uso de variables de entorno cargadas desde un fichero local secreto.

La solución elegida para gestionar variables de entorno deberá de ser fácil de usar y estar correctamente mantenida. En Go, las principales opciones son [Godotenv](https://github.com/joho/godotenv) y [Viper](https://github.com/spf13/viper).
#### 1. Viper
Viper es la opción de referencia cuando se trata de configuración en Go. Entre sus funciones, Viper es capaz no solo de gestionar ficheros de configuración para usar su contenido como variables de entorno sino que también de conectarse directamente con etcd y obtener la configuración del almacenamiento distribuido. Aunque usar Viper nos ahorraría la necesidad de usar dos herramientas distintas, su uso y configuración resultan algo más complejos y no he sido capaz de conseguir utilizarlo.
#### 2. Godotenv
Es la herramienta destinada a esta tarea más simple en Go. Cumple su función perfectamente y gracias a su paquete autoload no es necesario siquiera escribir una sola línea de código más allá del import para que cargue las variables de entorno de un fichero ".env". Es por su simplicidad, facilidad de uso y el buen estado de su repositorio de Github por lo que será la opción elegida.