package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	controllers "../../controllers"
	helpers "../../helpers"
)

//GamePageHandler : Entra al Juego
func GamePageHandler(response http.ResponseWriter, request *http.Request) {
	session, _ := store.Get(request, "cookie-name")

	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		http.Error(response, "Forbidden", http.StatusForbidden)
		return
	}

	userName := GetUserName(request)
	if !helpers.IsEmpty(userName) {
		cookieHash := request.Header.Get("Cookie")
		log.Println("CookieHASH :", cookieHash)

		worlController := controllers.WorldController{}
		worlController.Quit(userName)
		var indexBody, _ = helpers.LoadFile("template/index.html")
		indexPlayer := strings.Replace(indexBody, "${playerName}", userName, -1)

		fmt.Fprintf(response, indexPlayer)
	} else {
		http.Redirect(response, request, "/", 302)
	}
}
