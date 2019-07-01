package models

import (
	"time"

	structs "../models/structs"
	"gopkg.in/mgo.v2/bson"
)

// PlayerModel etc etc
type PlayerModel struct {
}

// InitLogin , etc etc
func (pm *PlayerModel) InitLogin(user string, cookieHash string, positionFree *structs.Object) bool {
	//Obtengo DAO de MongoDB
	isValid := true
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()

	if err == nil {
		player := structs.Player{}
		player.UserName = user
		player.CookieHash = cookieHash
		player.Update = time.Now()
		player.Position = positionFree

		playerDAO := session.DB(DataBaseName).C("Player")

		playerDAO.Upsert(
			bson.M{"username": player.UserName},
			bson.M{"$set": bson.M{"CookieHash": player.CookieHash, "Update": time.Now(), "Position": player.Position}},
		)
	} else {
		isValid = false
		panic(err)
	}
	return isValid
}
