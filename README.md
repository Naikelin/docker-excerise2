# Docker exercise 2

This exercise tries to imitate some situation that you may experience when building through the source code with Docker.

## Build

Prebuild: 
You will need: git and golang
 - On Debian based systems (Ubuntu, centos...):
> $ sudo apt-get -yqq update && sudo apt-get -yqq git golang

Building:
> $ git clone https://github.com/Naikelin/docker-exercise2  \
> $ cd docker-exercise2  \
> $ go build -o app server.go

Run:

> On folder directory:  \
> $ ./app

If you built a docker:

> $ docker run -it 'name of container' /bin/sh

#Español:

Ejercicio: 
Usted deberá buildear esta aplicación ejecutando el comando go build (go build construye un ejecutable para un archivo golang). Este ejecutable levantará un server en localhost, puerto 8080. La idea es que construya una imagen de docker funcional. Recuerde que para construir una imagen en docker deberá ejecutar en la carpeta donde tenga el Dockerfile (utilizando ubuntu como base):

> docker build -t exercise .

Una vez se haya construido la imagen, dependiendo de cómo se hizo, puede ejecutar la imagen de docker de distintas maneras.
Es recomendable que construya en el dockerfile el ejecutable y luego este sea añadido como entrypoint. Dado que el ejecutable levanta un servidor, usted puede realizar port-forward hacia su propia pc. En resumen, se le recomienda ejecutar de la siguiente forma el dockerfile:
> docker run -p 8080:8080 -it exercise

Usted al ejecutar lo enterior, obtendrá un output en su navegador ingresando a la ruta: http://localhost:8080. Si usted se encuentra en linux, también es posible que pueda ejecutar "wget http://localhost:8080" en su terminal. ¡Éxito!
