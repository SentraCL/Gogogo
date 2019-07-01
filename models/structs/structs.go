package structs

import "time"

const (
	//PlayerType , tipo Jugador
	PlayerType = 1
	//SentraPointType , tipo Moneda Sentra
	SentraPointType = 2
	//EnemyType , tipo Enemigo en el juego
	EnemyType = 3
)

//PlayerOnline : Acciones del jugador desde Front-End, tambien esta definida su equivalencia en JSON
type PlayerOnline struct {
	Name   string `json:"name"`
	Move   string `json:"move"`
	Sprite string `json:"sprite"`
}

//World :Definicion de posiciones en el juego
type World struct {
	Name    string
	MaxY    int
	MaxX    int
	Objects []Object
	Update  time.Time
}

//User :Usuario registrado en el juego
type User struct {
	UserName string
	Password string
	Email    string
	Create   time.Time
}

//Player :Entidad en el juego
type Player struct {
	UserName   string
	CookieHash string
	Score      int
	TopScore   int
	Position   *Object
	Update     time.Time
}

//Object : Objetos o actores en el juego
type Object struct {
	X, Y   int
	Full   bool
	Who    string
	Sprite string
	Type   int
}
