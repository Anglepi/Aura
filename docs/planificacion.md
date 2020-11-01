# Planificación del proyecto 

El proyecto no está planificado en su totalidad, pues varias funcionalidades aún no están totalmente aclaradas y habrá otros aspectos técnicos que aún desconozco y no puedo tener en cuenta actualmente. Aun así, he podido dividir el desarrollo inicial en 4 partes:

## Primera parte. Comenzar las bases de Aura

El desarrollo comenzará por esta parte, que consiste en únicamente comenzar el desarrollo del núcleo.

El núcleo se encargará de la comunicación del asistente con el usuario: será capaz de interpretar sus mensajes y recurrir a las funcionalidades necesarias para resolver las necesidades del usuario, pero como esto depende de otras muchas funcionalidades que aún no estarán implementadas, con recibir y devolver un saludo de momento será suficente, y eso será lo que se desarrolle en esta primera parte

Luego, durante el desarrollo del proyecto (nuevos componentes) se volverá a esta parte para incluir el uso de dichos componentes y a reconocer las peticiones de los usuarios que requieren éstos.

### Hitos e historias de usuario
 - [Comenzar la base de Aura](https://github.com/Anglepi/Aura/milestone/7)
     + [HU - Comunicación con el asistente](https://github.com/Anglepi/Aura/issues/9)

## Segunda parte. Algunas funcionalidades genéricas de todos los asistentes

La intención de incluir esta parte y en este preciso momento, es la de aprovechar el desarrollo de elementos ya conocidos, cuyos conceptos están claros, para coger experiencia con el lenguaje de programación y sus mecanismos y utilidades.

Se comenzará el desarrollo de esta parte con el componente de "Notas", que es simple y sencillo. Tras esta parte, comenzaré con el componente de "Recordatorios", que puede resultar muy similar al de las notas, pero con esa componente temporal.

Estas dos partes son muy simples, pues consisten en permitir insertar y consultar un texto en una base de datos

### Hitos e historias de usuario
 - [Algunas funciones genéricas de asistentes](https://github.com/Anglepi/Aura/milestone/8)
     + [HU - Notas](https://github.com/Anglepi/Aura/issues/12)
     + [HU - Recordatorios](https://github.com/Anglepi/Aura/issues/11)
     + [HU - Registro de usuario](https://github.com/Anglepi/Aura/issues/10)

## Tercera parte. Primeras funcionalidades atípicas

En esta tercera parte comenzaré el desarrollo de las funcionalidades que no son generalmente esperables dentro de los asistentes virtuales que existen en la actualidad. He hecho esta distinción con las funcionalidades genéricas ya que de momento, estos elementos atípicos que se me han ocurrido tienen algo más de complejidad.

En principio, la única funcionalidad que tengo pensada de esta parte es un gestor de memes, que te permita guardar y crear memes.

### Hitos e historias de usuario
 - [Primeras funciones atípicas](https://github.com/Anglepi/Aura/milestone/9)
     + [HU - Memes](https://github.com/Anglepi/Aura/issues/13)

## Parte paralela. Gestión de logs

Es muy probable que esta parte se acabe desarrollando paralelamente al resto, puesto que puede ser de ayuda a la hora de debuggear.

Esta parte consiste simplemente en implemenar [auralogger](https://github.com/Anglepi/Aura/blob/main/src/logs/auralogger.go) para que permita guardar información sobre el uso y correcto funcionamiento del sistema.

### Hitos e historias de usuario
 - [Gestión de logs](https://github.com/Anglepi/Aura/milestone/6)
     + [HU - Log de acciones sobre los datos](https://github.com/Anglepi/Aura/issues/8)
     + [HU - Registro de logs](https://github.com/Anglepi/Aura/issues/7)
     + [HU - Almacenar información](https://github.com/Anglepi/Aura/issues/15)