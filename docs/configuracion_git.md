# Configuración de git

## Introducción

A continuación se expondrá una guía sobre cómo se ha configurado todo lo relacionado a git: creación de clave privada y pública y subida de ésta a GitHub, configuración del nombre, email y foto de perfil de la cuenta para que se muestre correctamente en los commits y activación de un segundo factor de autenticación. Todo esto hecho en un sistema Ubuntu.

## Creación de clave pública y privada

Para realizar esta tarea, se ha seguido la [documentación oficial de GitHub](https://docs.github.com/es/free-pro-team@latest/github/authenticating-to-github/generating-a-new-ssh-key-and-adding-it-to-the-ssh-agent), es una tarea mucho más sencilla de lo que puede parecer.

En primer lugar, voy a crear crear la clave. Para ello abro un terminal y escribo lo siguiente, indicando mi direccion de correo:

> ssh-keygen -t rsa -b 4096 -C angle@correo.ugr.es

Esto comenzará un proceso de generación de clave, y nos preguntará en que ruta queremos almacenarla. En mi caso, yo escogí la siguiente ruta:

> /home/angel/.ssh/DOCS_clave_publica

De esta manera, me almacenará dentro del directorio .shh un fichero llamado DOCS_clave_publica (el nombre y la ruta que usé no tienen por qué ser los expuestos en esta documentación).

A continuación solicitará una "paraphrase" como mecanismo de seguridad adicional, que protege todos los sitios en los que usemos dicha clave si alguien lograra acceder a mi equipo. Se puede dejar en blanco si no se desea usar este mecaniso.

Inicio el agente de ssh en caso de no tenerlo iniciado.

> eval "$(ssh-agent -s)"

Y añado la clave privada al ssh

> ssh-add /home/angel/.ssh/DOCS_clave_publica

Lo único que queda por hacer es subir la clave pública a GitHub. Para ello, abro la configuración de mi perfil, y en la sección "SSH and GPG keys" le doy al botón de "New SSH Key".

Para añadir la clave pública, primero tengo que establecerle un título, y luego añadir el contenido de /home/angel/.ssh/DOCS_clave_publica.pub (importante el .pub). Una vez hecho, compruebo que funciona correctamente.

> ssh -T git@github.com

Esto nos responderá con un mensaje que indica si la comunicación es correcta y estás autenticado, o si ha ocurrido algún problema.


## Configuración del perfil

Hay que establecer un nombre y un email adecuados para que se muestre a la hora de hacer commits.

> git config --global user.name "Ángel Píñar Rivas"
> git config --global user.email angle@correo.ugr.es

Se debe tener en cuenta que con "--global" quiero indicar que se cambie en la configuración global de mi perfil, y que si quisiera que esto solo ocurriese para un repositorio concreto, tendría que navegar al directorio del repositorio y ejecutar las líneas en cuestion sin el tag "--global".

## Activación del segundo factor de autenticación

De nuevo, en la sección de configuración de mi perfil de GitHub, en el apartado de "Account Security" puedes elegir diversos métodos como segundo factor de autenticación. En mi caso, he elegido el de SMS a un número de teléfono, ya que la otra opción me obliga a instalar la aplicación móvil y no creo que me aporte mucho tener dicha aplicación instalada, ya que a mí personalmente me gusta usar el teléfono únicamente como herramienta de comunicación.