package controllers

// SessionController posee los atributos y los metodos para el manejo de login
type SessionController struct {
	players []string
}

// DoLogin , Valida usuario y posterior a esto registra su inicio de sesion
// Notificar a Login Controller que Inicio una nueva sesion
// Agregar conexion a arreglo del socket
func (sc SessionController) DoLogin(user, password, cookieHash string) bool {
	isValid := false
	if sesionModel.IsUserValid(user, password) {
		//Registrar Inicio de Sesion
		player := worldController.GetFreePosition(user)

		playerModel.InitLogin(user, cookieHash, player)
		//Agrega el jugador al listado de jugadores online, via socket
		isValid = true
	}
	//Ojo que solo valida que el usuario no ingreso bien las credenciales
	return isValid
}

//Exit ,Adios mundo Cruel
func (sc SessionController) Exit(user, cookieHash string) bool {
	isValid := worldController.Quit(user)
	return isValid
}

//CreateUser :crea usuario si este no se encuentra en la BD
func (sc SessionController) CreateUser(user, email, password string) bool {

	if sesionModel.CreateUser(user, email, password) {
		return true
	}

	return false
}
