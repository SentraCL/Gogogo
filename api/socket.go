package api

import (
	"encoding/json"
	"log"

	controllers "../controllers"
	helpers "../helpers"
	structs "../models/structs"
	socketio "github.com/googollee/go-socket.io"
)

//SocketGame socket del juego!
type SocketGame struct {
}

//GetServer Servidor Socket de GOGOGO
func (sg SocketGame) GetServer() *socketio.Server {
	worlController := controllers.WorldController{}

	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.Println("connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "move", func(s socketio.Conn, playerJsonString string) string {
		log.Println("connected:", s.ID())
		player := structs.PlayerOnline{}
		json.Unmarshal([]byte(playerJsonString), &player)
		playerCtrl := controllers.PlayerController{}
		playerCtrl.MovePlayer(&player)
		world := worlController.GetWorld()
		worldString := helpers.StringifyJSON(world)
		//Contexto comun de socket para poder renderizar en el front-end
		s.SetContext("")
		s.Emit("renderWorld", worldString)
		return "El jugador " + player.Name + " se movio hacia :" + player.Move
	})

	server.OnEvent("/", "updateWorld", func(s socketio.Conn, localWorld string) {
		world := worlController.GetWorld()
		worldString := helpers.StringifyJSON(world)
		//Contexto comun de socket para poder renderizar en el front-end

		s.SetContext("")
		if localWorld != worldString {
			log.Println("connected:", s.ID())
			log.Println("UpdateWorld:", localWorld)
			s.Emit("renderWorld", worldString)
		}

	})

	server.OnEvent("/", "bye", func(s socketio.Conn) string {
		last := s.Context().(string)
		s.Emit("bye", last)
		s.Close()
		return last
	})

	server.OnError("/", func(e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		log.Println("closed", msg)
	})

	return server
}
