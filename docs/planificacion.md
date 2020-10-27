# Planificación del proyecto 

Actualmente he dividido el desarrollo del proyecto en 4 partes, una de las cuales se desarrollará de forma paralela a las otras 3.

## Primera parte. Núcleo de Aura

El desarrollo comenzará por esta parte, que consiste en desarrollar un nucleo central, en el que se le da personalidad al asistente.

El núcleo se encargará de la comunicación del asistente con el usuario: será capaz de interpretar sus mensajes y recurrir a las funcionalidades necesarias para resolver las necesidades del usuario.

En primer lugar, se tiene como objetivo construir una interfaz básica del núcleo, que sea capaz de recibir mensajes y devolver respuestas. Luego, durante el desarrollo del proyecto (nuevos componentes) se volverá a este núcleo para enseñar a Aura a usar dichos componentes y a reconocer las peticiones de los usuarios que requieren éstos.

### Hitos e historias de usuario
 - [Núcleo de Aura](https://github.com/Anglepi/Aura/milestone/7)
     + [HU - Comunicación con el asistente](https://github.com/Anglepi/Aura/issues/9)

## Segunda parte. Funcionalidades genéricas de todos los asistentes

La intención de incluir esta parte y en este preciso momento, es la de aprovechar el desarrollo de elementos ya conocidos, cuyos conceptos están claros, para coger experiencia con el lenguaje de programación y sus mecanismos y utilidades.

Se comenzará el desarrollo de esta parte con el componente de "Notas", que es simple y sencillo. Tras esta parte, comenzaré con el componente de "Recordatorios", que puede resultar muy similar al de las notas, pero con esa componente temporal.

### Hitos e historias de usuario
 - [Funciones genéricas de asistentes](https://github.com/Anglepi/Aura/milestone/8)
     + [HU - Notas](https://github.com/Anglepi/Aura/issues/12)
     + [HU - Recordatorios](https://github.com/Anglepi/Aura/issues/11)
     + [HU - Registro de usuario](https://github.com/Anglepi/Aura/issues/10)

## Tercera parte. Funcionalidades atípicas

Aquí ya comenzaré el desarrollo de las funcionalidades que no son generalmente esperables dentro de los asistentes virtuales que existen en la actualidad. He hecho esta distinción con las funcionalidades genéricas ya que de momento, estos elementos atípicos que se me han ocurrido tienen algo más de complejidad.

En principio, la única funcionalidad que tengo pensada de esta parte es un gestor de memes, que te permita guardar y crear memes.

### Hitos e historias de usuario
 - [Funciones atípicas](https://github.com/Anglepi/Aura/milestone/9)
     + [HU - Memes](https://github.com/Anglepi/Aura/issues/13)

## Parte paralela. Gestión de logs

Durante el desarrollo de las tres partes anteriores, se irá también trabajando en la parte de logs, ya que ayudará en el desarrollo del proyecto.

Se incluirán logs tanto para registrar el comportamiento del programa como para registrar las acciones que realizan los usuarios.

### Hitos e historias de usuario
 - [Gestión de logs](https://github.com/Anglepi/Aura/milestone/6)
     + [HU - Log de acciones sobre los datos](https://github.com/Anglepi/Aura/issues/8)
     + [HU - Registro de logs](https://github.com/Anglepi/Aura/issues/7)
     + [HU - Almacenar información](https://github.com/Anglepi/Aura/issues/15)