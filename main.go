package main

import (
	"flag"
	"log"
	"net/http"
	"strings"

	Api "./api"
	Routes "./routes"
)

func main() {
	var port string
	var enemies int
	flag.StringVar(&port, "p", ":8080", "Puerto en el cual correra la aplicacion")
	flag.IntVar(&enemies, "e", 1, "Cantidad de Enemigos")
	flag.Parse()

	r := Routes.Route{}
	api := Api.SocketGame{}
	game := Api.GameManager{}
	socket := api.GetServer()
	go socket.Serve()
	go game.Start(enemies)

	defer socket.Close()

	http.Handle("/socket.io/", socket)
	http.Handle("/assets/",
		http.StripPrefix(strings.TrimRight("/assets/", "/"),
			http.FileServer(http.Dir("./template/assets"))))
	http.Handle("/", r.Routers())
	log.Fatal(http.ListenAndServe(":"+port, nil))
	log.Println("Let's GOGOGOGOGOGO!!")

}
