package api

import (
	"log"
	"time"

	controllers "../controllers"
	structs "../models/structs"
)

//GameManager , La mente detras del juego
type GameManager struct {
}

//SeePlayer , Rango de Vista del Enemigo
func (gm GameManager) SeePlayer(i, min, max int) bool {
	return (i > min-1) && (i < max+1)
}

//Start , Iniciar Juego!!!
func (gm GameManager) Start(enemiesMax int) {

	log.Println("Que empiece el Juego ")
	go gm.Enemies(enemiesMax)
}

//Enemies , Mente Enemiga!!
func (gm GameManager) Enemies(enemiesMax int) {
	log.Println("Enemigos Listos!!", enemiesMax)
	worlController := controllers.WorldController{}
	enemies := worlController.GetEnemies()

	//sEnemy := helpers.StringifyJSON(enemies)
	if len(*enemies) == 0 || len(*enemies) < enemiesMax {
		worlController := controllers.WorldController{}
		for e := 0; e < enemiesMax; e++ {
			log.Println("Crear Enemigo!")
			worlController.CreateEnemy()
		}
	}

	for true {
		enemies = worlController.GetEnemies()
		//Obtener listado de Enemigos
		players := *worlController.GetPlayers()
		//Itera Enemigos y les da Algo Parecido a la Inteligencia ( :3 instinto mas que nada )

		for _, enemy := range *enemies {
			othersEnemies := []structs.Object{}
			for _, otherEnemy := range *enemies {
				if otherEnemy.Who != enemy.Who {
					othersEnemies = append(othersEnemies, otherEnemy)
				}
			}

			go gm.EnemyGetIA(&enemy, players, othersEnemies)
			time.Sleep(1 * time.Second)
		}

		//En el caso que eliminen a los Enemigos!
		if len(*enemies) == 0 {
			worlController := controllers.WorldController{}
			worlController.CreateEnemy()
		}
	}
}

//EnemyGetIA : Instinto Artificial para Enemigo
func (gm GameManager) EnemyGetIA(enemy *structs.Object, players []structs.Object, others []structs.Object) {
	goToX := 0
	goToY := 0
	victim := ""
	//Existe un Jugador dentro del Rango de Vista del Enemigo
	for _, player := range players {
		if player.Type == structs.PlayerType {
			//Busca a 3 Niveles de rango que jugador esta mas cerca
			for l := 1; l <= 6; l++ {
				maxX := enemy.X + (l * 2)
				minX := enemy.X - (l * 2)
				maxY := enemy.Y + (l * 2)
				minY := enemy.Y - (l * 2)
				//Guarda el Jugador que este mas cerca

				if gm.SeePlayer(player.X, minX, maxX) && gm.SeePlayer(player.Y, minY, maxY) {
					goToX = player.X
					goToY = player.Y
					victim = player.Who
				}
			}
		}
	}
	if victim != "" {
		log.Println(enemy.Who + " se quiere comer a " + victim)
		collision := false
		//Que no choque con otro Enemigo

		worlController := controllers.WorldController{}
		if enemy.X > goToX {
			enemy.X--
		} else if enemy.X < goToX {
			enemy.X++
		}

		if enemy.Y > goToY {
			enemy.Y--
		} else if enemy.Y < goToY {
			enemy.Y++
		}

		for _, other := range others {
			if enemy.X == other.X && enemy.Y == other.Y {
				collision = true
			}
		}

		if enemy.X == goToX && enemy.Y == goToY {
			worlController.Quit(victim)
			enemy.Who = victim
		}
		//Si no colisiono, que busque al jugador
		if !collision {
			worlController.MoveObject(enemy)
		}

	}

}
