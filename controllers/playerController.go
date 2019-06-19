package controllers

import (
	strutcs "../models/structs"
)

//PlayerController COntrolador de los jugadores
type PlayerController struct {
}

//MovePlayer mueve el player
func (pc PlayerController) MovePlayer(playerMove *strutcs.PlayerOnline) {

	maxX := 10
	maxY := 20

	playerPos := worldController.GetPlayerPosition(playerMove.Name)

	if playerMove.Move == "Down" {
		playerPos.X++
	}

	if playerMove.Move == "Up" {
		playerPos.X--
	}

	if playerMove.Move == "Left" {
		playerPos.Y--
	}
	if playerMove.Move == "Right" {
		playerPos.Y++
	}

	if playerPos.Y >= maxY+1 {
		playerPos.Y = 1
	}
	if playerPos.X >= maxX+1 {
		playerPos.X = 1
	}

	if playerPos.Y <= 0 {
		playerPos.Y = maxY
	}
	if playerPos.X <= 0 {
		playerPos.X = maxX
	}
	who := ""
	isFreePosition := worldModel.GetWhoAreIn(playerPos.X, playerPos.Y, &who)
	if isFreePosition {
		worldModel.RemoveInTheWorld(playerMove.Name)
		worldModel.PutInTheWorld(playerPos.X, playerPos.Y, playerMove.Name)
	}
}
