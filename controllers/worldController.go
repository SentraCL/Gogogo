package controllers

import (
	"log"
	"math/rand"
	"strconv"

	structs "../models/structs"
)

// WorldController Representacion del mundo del Juego
type WorldController struct {
	name string
}

//CreateEnemy : Crea un Enemigo
func (sc WorldController) CreateEnemy() *structs.Object {
	enemy := structs.Object{}

	nBot := strconv.Itoa(rand.Intn(1000))
	enemy.Who = "Bot #" + nBot
	enemy.Type = structs.EnemyType
	//TODO : Obtener valores minimos y maximos desde la base de datos segun el mundo.
	min := 1
	max := 15
	x := 0
	y := 0

	isFreePosition := false
	for !isFreePosition {
		x = rand.Intn(max-min) + min
		y = rand.Intn(max-min) + min
		enemy.X = x
		enemy.Y = y
		isFreePosition = worldModel.GetWhoAreIn(x, y, &enemy.Who)
	}

	enemy = *worldModel.CreateObject(enemy)
	return &enemy
}

//GetEnemies : Traeme los Enemigos!!
func (sc WorldController) GetEnemies() *[]structs.Object {
	enemies := worldModel.GetObjectByType(structs.EnemyType)
	return enemies
}

//GetPlayers : Traeme los Jugadores.
func (sc WorldController) GetPlayers() *[]structs.Object {
	players := worldModel.GetObjectByType(structs.PlayerType)
	return players
}

//GetPlayerPosition : Posicion del Jugador
func (sc WorldController) GetPlayerPosition(player string) *structs.Object {
	playerPosition := worldModel.GetPlayer(player)
	return playerPosition
}

//GetWorld : Posicion del Jugador
func (sc WorldController) GetWorld() *structs.World {
	world := worldModel.Get()
	return world
}

//GetFreePosition :Retorna una posicion sin uso en el mundo
func (sc WorldController) GetFreePosition(player string) *structs.Object {
	freePosition := structs.Object{}
	freePosition.Who = player

	min := 1
	max := 15
	x := 0
	y := 0

	isFreePosition := false
	for !isFreePosition {
		x = rand.Intn(max-min) + min
		y = rand.Intn(max-min) + min

		freePosition.X = x
		freePosition.Y = y

		isFreePosition = worldModel.GetWhoAreIn(x, y, &player)
		if !isFreePosition {
			//log.Printf("Quien esta en la posicion : %s \n" + player)
		} else {
			//log.Printf("Posicion Libre para %s \n", player)
			worldModel.PutInTheWorld(x, y, player)
		}
	}
	return &freePosition
}

//Quit :Retorna una posicion sin uso en el mundo
func (sc WorldController) Quit(player string) bool {

	isOut := worldModel.RemoveInTheWorld(player)
	if !isOut {
		log.Println("Error no quiere irse del mundo")
	}

	return isOut
}

//MoveObject :Quita un objecto de su ubicacion anterior y lo coloca en la seteada por parametros
func (sc WorldController) MoveObject(object *structs.Object) bool {
	Who := ""
	isOut := worldModel.GetWhoAreIn(object.X, object.Y, &Who)
	if isOut {
		isOut = worldModel.MoveInTheWorld(object)
	}
	return isOut
}
