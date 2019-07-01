package models

import (
	"log"
	"time"

	strutcs "../models/structs"
	"gopkg.in/mgo.v2/bson"
)

// SessionModel etc etc
type SessionModel struct {
}

//IsUserValid , valida crenciales
func (sm *SessionModel) IsUserValid(user, pass string) bool {
	//Obtener Conexion.
	session, err := GetSession()
	//Se ejecuta una vez salido de la funcion
	defer session.Close()

	userDAO := session.DB(DataBaseName).C("User")
	userResult := strutcs.User{}
	err = userDAO.Find(bson.M{"username": user, "password": pass}).One(&userResult)
	_user, _pass := userResult.UserName, userResult.Password
	isValid := false

	if err == nil {
		isValid = (user == _user && pass == _pass)
	} else {
		isValid = false
	}
	log.Println("Es un Usuario Valido : ", isValid)
	return isValid
}

//CreateUser crea usuarios si no estan registrados
func (sm *SessionModel) CreateUser(user, email, pass string) bool {
	isCreated := false
	//Obtener Conexion.
	session, err := GetSession()
	defer session.Close()

	userDB := session.DB(DataBaseName).C("User")

	userType := strutcs.User{}
	//user, pass, email
	userType.UserName = user
	userType.Password = pass
	userType.Email = email
	userType.Create = time.Now()

	err = userDB.Insert(userType)
	if err != nil {
		log.Fatal(err)
	} else {
		isCreated = true
	}

	result := strutcs.User{}
	err = userDB.Find(bson.M{"username": user}).One(&result)
	return isCreated
}
