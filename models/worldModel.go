package models

import (
	"log"

	structs "../models/structs"
	"gopkg.in/mgo.v2/bson"
)

// WorldModel etc etc
type WorldModel struct {
}

// Get Obten el Mundo
func (wml WorldModel) Get() *structs.World {

	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	whoInWorld := structs.World{}

	worldDAO := session.DB(DataBaseName).C("World")
	err = worldDAO.Find(bson.M{"name": "Mundo1"}).One(&whoInWorld)

	//log.Printf("who %s \n\r", helpers.StringifyJSON(&whoInWorld))

	return &whoInWorld
}

// GetWhoAreIn , etc etc
func (wml WorldModel) GetWhoAreIn(x, y int, whois *string) bool {
	//Obtengo DAO de MongoDB
	isValid := true
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	whoInWorld := structs.Object{}
	worldDAO := session.DB(DataBaseName).C("World")
	err = worldDAO.Find(bson.M{"name": "Mundo1"}).Select(bson.M{"x": x, "y": y}).One(&whoInWorld)

	//log.Printf("who %s \n\r", helpers.StringifyJSON(&whoInWorld))

	isValid = !whoInWorld.Full
	whois = &whoInWorld.Who
	if isValid {
		//log.Println("Hay espacio libre!!")
	}
	return isValid
}

//PutInTheWorld ,
func (wml WorldModel) PutInTheWorld(x, y int, playerName string) bool {
	isValid := true
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	worldDAO := session.DB(DataBaseName).C("World")
	objects := []structs.Object{}

	worldType := structs.World{}

	err = worldDAO.Find(bson.M{"name": "Mundo1"}).One(&worldType)
	objects = worldType.Objects
	playerObj := structs.Object{}
	playerObj.Type = structs.PlayerType
	playerObj.X = x
	playerObj.Y = y
	playerObj.Who = playerName
	playerObj.Full = true
	objects = append(objects, playerObj)

	worldDAO.Upsert(
		bson.M{"name": "Mundo1"},
		bson.M{"$set": bson.M{"objects": objects}},
	)

	return isValid
}

//MoveInTheWorld , mueve un objecto en el mundo
func (wml WorldModel) MoveInTheWorld(object *structs.Object) bool {
	isValid := true

	isOut := wml.RemoveInTheWorld(object.Who)
	if !isOut {
		log.Println("Error no quiere irse del mundo")
	}

	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	worldDAO := session.DB(DataBaseName).C("World")
	worldType := structs.World{}
	err = worldDAO.Find(bson.M{"name": "Mundo1"}).One(&worldType)
	objects := worldType.Objects
	objects = append(objects, *object)

	worldDAO.Upsert(
		bson.M{"name": "Mundo1"},
		bson.M{"$set": bson.M{"objects": objects}},
	)

	return isValid
}

//RemoveInTheWorld ,
func (wml WorldModel) RemoveInTheWorld(who string) bool {
	//Obtengo DAO de MongoDB
	isValid := true
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	worldDAO := session.DB(DataBaseName).C("World")
	worldType := structs.World{}
	err = worldDAO.Find(bson.M{"name": "Mundo1"}).One(&worldType)
	objects := worldType.Objects

	for i := len(objects) - 1; i >= 0; i-- {
		object := objects[i]
		// Te encontre , te expulso
		if object.Who == who {
			//Agrego todo menos el punto donde me encuentro
			objects = append(objects[:i], objects[i+1:]...)
		}
	}

	worldDAO.Upsert(
		bson.M{"name": "Mundo1"},
		bson.M{"$set": bson.M{"objects": objects}},
	)

	return isValid
}

//GetPlayer ,
func (wml WorldModel) GetPlayer(playerName string) *structs.Object { //Obtengo DAO de MongoDB
	playerPos := structs.Object{}
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()

	if err != nil {
		panic(err)
	}

	worldDAO := session.DB(DataBaseName).C("World")
	objects := []structs.Object{}

	worldType := structs.World{}

	err = worldDAO.Find(bson.M{"name": "Mundo1"}).One(&worldType)
	objects = worldType.Objects

	for i := len(objects) - 1; i >= 0; i-- {
		object := objects[i]
		// Te encontre , te expulso
		if object.Who == playerName && object.Type == structs.PlayerType {
			playerPos = object
		}
	}

	return &playerPos
}

//GetObjectByType , obtiene segun tipo
func (wml WorldModel) GetObjectByType(typeObject int) *[]structs.Object { //Obtengo DAO de MongoDB
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()

	if err != nil {
		panic(err)
	}

	worldDAO := session.DB(DataBaseName).C("World")
	objects := []structs.Object{}
	worldType := structs.World{}
	err = worldDAO.Find(bson.M{"name": "Mundo1"}).One(&worldType)
	objects = worldType.Objects

	filterObject := []structs.Object{}

	for i := len(objects) - 1; i >= 0; i-- {
		object := objects[i]
		if object.Type == typeObject {
			filterObject = append(filterObject, object)
		}
	}
	return &filterObject
}

//CreateObject , obtiene segun tipo
func (wml WorldModel) CreateObject(object structs.Object) *structs.Object { //Obtengo DAO de MongoDB
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	worldDAO := session.DB(DataBaseName).C("World")
	objects := []structs.Object{}

	worldType := structs.World{}

	err = worldDAO.Find(bson.M{"name": "Mundo1"}).One(&worldType)
	objects = worldType.Objects
	objects = append(objects, object)

	worldDAO.Upsert(
		bson.M{"name": "Mundo1"},
		bson.M{"$set": bson.M{"objects": objects}},
	)

	return &object
}
