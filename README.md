![gogogo](https://github.com/SentraCL/Gogogo/blob/master/template/assets/img/playerUp.gif)
# Gogogo
Aplicacion de referencia de GO, contiene Hilo, Persistencia MongoDB, Lectura de Archivos, Manejo de Sesion y Socket.

Propone una estructura de proyecto separada por patron MVC.

# Prerequisitos

Tener instalado mongodb, antes de ejecutar la aplicacion (mongodb -dbpath ./db)

# Ejecucion del juego

Para ejecutar el juego, se deben setear los parametros p = puerto y e= cantidad de enemigos..

go run main.go -p 8080 -e 1

# Carpetas

api => Se encuentra la Rutina de Socket y de Manajo de Enemigos.

controllers => Contiene los Controllers segun.

db => Donde reside la base de datos de mongodb.

helpers => Utilitarios del Juego.

models => Representacion de la Base de Datos.

routes => Navegacion de la Aplicacion.

template => Template HTML.
