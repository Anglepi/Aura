# AURA
> Asistente Ubicuo Realmente Atípico.

Este repositorio contiene todo lo relacionado con un proyecto ideado por mi parte que prentende solventar problemas que me suelo encontrar en mi vida personal, los cuales de algún modo me obligan a realizar algún mínimo esfuerzo y me han hecho pensar que son fácilmente resolvibles mediante el uso del software.

## Arquitectura del sistema

Para resolver este sistema, voy a recurrir a una arquitectura dirigida por eventos, ya que va a depender de los eventos que genere el usuario para actuar.

Una primera aproximación a la estructuración del sistema que planteo sería la siguiente: un componente central, al que daré el nombre del proyecto, que tendrá una cola de mensajes, siendo estos las peticiones de todos los usuarios junto con parámetros que les identifiquen.

Por cada mensaje recibido, Aura lo procesará y determinará en función de éste que otro componente es necesario utilizar: uno matemático, otro de calendario, de recordatorios... y generará un evento dirigido a dicho componente, para que éste lo procese y devuelva una respuesta.

Tras recibir la respuesta, Aura le dará formato en lenguaje natural y contestará al mensaje que el usuario envió en primer lugar.

## Lenguaje a utilizar en el proyecto

He estado echando un vistazo a los distintos lenguajes de programación que nunca había utilizado. En principio me planteé Scala y Perl, acabé descartando Scala porque sentía que no encajaba del todo con un proyecto de esta naturaleza (desde mi desconocimiento del lenguaje), y poco despues un amigo me mencionó el lenguaje Go y me dio buenas opiniones, se ofreció a enseñarme algunos sitios web con ejemplos y me acabó gustando más que Perl.

Así que por lo tanto, el lenguaje que usaré para desarrollar este sistema será Go.

## Documentación del proyecto

+ [Puesta a punto del repositorio y configuración de git.](https://github.com/Anglepi/Aura/blob/main/docs/configuracion_git.md)
+ [Descripción del problema que se propone resolver, tema correspondiente al hito 0](https://github.com/Anglepi/Aura/blob/main/docs/hitos/hito0.md)