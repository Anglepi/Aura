# AURA
> Asistente Ubicuo Realmente Atípico.

Este repositorio contiene todo lo relacionado con un proyecto ideado por mi parte que pretende solventar problemas que me suelo encontrar en mi vida personal, los cuales de algún modo me obligan a realizar algún mínimo esfuerzo y me han hecho pensar que son fácilmente resolubles mediante el uso del software.

Consiste, como indica el título, en un asistente virtual que será capaz de interpretar las peticiones de un usuario y responder de forma adecuada.

## Arquitectura del sistema

En este proyecto se plantea usar una **arquitectura basada en componentes**.

Debido a la naturaleza multifuncional del sistema, considero que lo más adecuado es dividir estas distintas funcionalidades en componentes. Uno de los elementos del sistema, el principal, será [aura](https://github.com/Anglepi/Aura/blob/main/src/aura.go), que se encarga de interpretar las peticiones del usuario, y decidir que componente es el que debe utilizar para cumplir su petición y elaborar la respuesta.

Cada uno de los componentes aportará una funcionalidad única al sistema. Una primera aproximación a los diferentes componentes del sistema incluyen la toma de notas, uso de recordatorios y gestión de memes para resolver peticiones del usuario, y gestión de base de datos, usuarios y logs para uso interno del sistema.

Se ha elaborado un código esqueleto que incluye los componentes mencionados anteriormente. En la sección de [clases utilizadas](#clases-utilizadas) se encuentran los enlaces a los archivos fuente.

## Lenguaje a utilizar en el proyecto

He estado echando un vistazo a los distintos lenguajes de programación que nunca había utilizado. En principio me planteé Scala y Perl, acabé descartando Scala porque sentía que no encajaba del todo con un proyecto de esta naturaleza (desde mi desconocimiento del lenguaje), y poco después un amigo me mencionó el lenguaje Go y me dio buenas opiniones, se ofreció a enseñarme algunos sitios web con ejemplos y me acabó gustando más que Perl.

Así que por lo tanto, el lenguaje que usaré para desarrollar este sistema será Go.

## Planificación

He elaborado una planificación para el proyecto, aunque algo genérica y probablemente incompleta, sujeta a cambios en las próximas semanas.

Esta falta de precisión en la planificación se debe principalmente a mi inexperiencia con el lenguaje y desconocimiento actual sobre las distintas herramientas y mecanismos que necesitaré aprender para intercomunicar los componentes, habilitar un mecanismo adecuado para la comunicación del agente, desplegar el proyecto...

La planificación se puede encontrar en [este enlace](https://github.com/Anglepi/Aura/blob/main/docs/planificacion.md)

## Clases utilizadas

Esta es la estructura de clases sobre la que estoy trabajando para desarrollar este proyecto:

 + [src/aura.go](https://github.com/Anglepi/Aura/blob/main/src/aura.go)
 + src/auramemory/
     * [auramemory.go](https://github.com/Anglepi/Aura/blob/main/src/auramemory/auramemory.go)
 + src/logs/
     * [auralogger.go](https://github.com/Anglepi/Aura/blob/main/src/logs/auralogger.go)
 + src/memes/
     * [meme.go](https://github.com/Anglepi/Aura/blob/main/src/memes/meme.go)
     * [mememanager.go](https://github.com/Anglepi/Aura/blob/main/src/memes/mememanager.go)
 + src/notes/
     * [note.go](https://github.com/Anglepi/Aura/blob/main/src/notes/note.go)
     * [notemanager.go](https://github.com/Anglepi/Aura/blob/main/src/notes/notemanager.go)
 + src/reminders/
     * [reminder.go](https://github.com/Anglepi/Aura/blob/main/src/reminders/reminder.go)
     * [remindermanager.go](https://github.com/Anglepi/Aura/blob/main/src/reminders/remindermanager.go)
 + src/users
     * [user.go](https://github.com/Anglepi/Aura/blob/main/src/users/user.go)
     * [usermanager.go](https://github.com/Anglepi/Aura/blob/main/src/users/usermanager.go)

## Documentación del proyecto

+ [Puesta a punto del repositorio y configuración de git.](https://github.com/Anglepi/Aura/blob/main/docs/configuracion_git.md)
+ [Descripción del problema que se propone resolver](https://github.com/Anglepi/Aura/blob/main/docs/milestones/milestone0.md)
+ [Planificación del proyecto](https://github.com/Anglepi/Aura/blob/main/docs/planificacion.md)