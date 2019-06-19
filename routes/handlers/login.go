package handlers

import (
	"fmt"
	"log"
	"net/http"

	controllers "../../controllers"
	helpers "../../helpers"
	"github.com/gorilla/securecookie"
)

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))

var sesionCtrl = controllers.SessionController{}

// Handlers

//LoginPageHandler : Muestra la pagina de login
func LoginPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("template/login.html")
	fmt.Fprintf(response, body)
}

//LoginHandler : registra el login del usuario
func LoginHandler(response http.ResponseWriter, request *http.Request) {
	name := request.FormValue("name")
	pass := request.FormValue("password")
	redirectTarget := "/"
	if !helpers.IsEmpty(name) && !helpers.IsEmpty(pass) {
		cookieHash := UpSetCookie(name, response)
		//cookieHash := request.Header.Get("Cookie")
		log.Println("Inicio Sesion :", name)
		log.Println("CookieHASH :", cookieHash)
		isLogin := sesionCtrl.DoLogin(name, pass, cookieHash)

		if isLogin {
			redirectTarget = "/index"
		} else {
			redirectTarget = "/?Error"
		}
	}
	http.Redirect(response, request, redirectTarget, 302)
}

//RegisterPageHandler : Imprime formulario de Registro
func RegisterPageHandler(response http.ResponseWriter, request *http.Request) {
	var body, _ = helpers.LoadFile("template/registro.html")
	fmt.Fprintf(response, body)
}

//RegisterHandler : Guarda el nuevo registro de Usuario
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	userName := r.FormValue("username")
	email := r.FormValue("email")
	password := r.FormValue("password")
	confirmpass := r.FormValue("confirmPassword")

	_userName, _email, _pass, _confirmpass := false, false, false, false
	_userName = !helpers.IsEmpty(userName)
	_email = !helpers.IsEmpty(email)
	_pass = !helpers.IsEmpty(password)
	_confirmpass = !helpers.IsEmpty(confirmpass)

	if _userName && _email && _pass && _confirmpass {
		sesionCtrl.CreateUser(userName, email, password)
		var indexBody, _ = helpers.LoadFile("template/login.html")
		fmt.Fprintf(w, indexBody)
	} else {
		fmt.Fprintln(w, "Falto llenar un campo!")
	}
}

// LogoutHandler : Cierra Sesion
func LogoutHandler(response http.ResponseWriter, request *http.Request) {
	//quitar conexion a arreglo del socket
	userName := GetUserName(request)
	cookieHash := request.Header.Get("Cookie")
	sesionCtrl.Exit(userName, cookieHash)
	ClearCookie(response)
	worlController := controllers.WorldController{}
	defer worlController.Quit(userName)
	http.Redirect(response, request, "/", 302)
}

//UpSetCookie : Setea nombre de jugador en cookie y retorna la cookie encriptada
func UpSetCookie(userName string, response http.ResponseWriter) string {
	cookieCode := ""

	value := map[string]string{
		"name": userName,
	}
	if encoded, err := cookieHandler.Encode("cookie", value); err == nil {
		cookie := &http.Cookie{
			Name:  "cookie",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(response, cookie)
		cookieCode = encoded
	}

	return cookieCode
}

//ClearCookie : Limpia las cookies
func ClearCookie(response http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:   "cookie",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(response, cookie)
}

//GetUserName : Obtiene nombre del jugador
func GetUserName(request *http.Request) (userName string) {
	if cookie, err := request.Cookie("cookie"); err == nil {
		cookieValue := make(map[string]string)
		if err = cookieHandler.Decode("cookie", cookie.Value, &cookieValue); err == nil {
			userName = cookieValue["name"]
		}
	}
	return userName
}
