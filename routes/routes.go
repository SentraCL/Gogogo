package routes

import (
	handlers "./handlers"
	"github.com/gorilla/mux"
)

// Route listados de rutas
type Route struct {
}

// Routers , retorna los routes asociados a la rutina de login
func (r Route) Routers() *mux.Router {
	route := mux.NewRouter()
	route.HandleFunc("/", handlers.LoginPageHandler)     // GET
	route.HandleFunc("/index", handlers.GamePageHandler) // GET
	route.HandleFunc("/login", handlers.LoginHandler).Methods("POST")
	route.HandleFunc("/register", handlers.RegisterPageHandler).Methods("GET")
	route.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")
	route.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")
	return route
}
