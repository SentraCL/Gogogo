package models

import "gopkg.in/mgo.v2"

//DataBaseName :Nombre de la Base de Datos
const DataBaseName = "GOGOGO"

//GetSession , Obtener session
func GetSession() (*mgo.Session, error) {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return session, err
}
