package models

import (
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

type Poll struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Question    string            `json:"question"`
	Options     []string          `json:"options"`
	Votes       map[string]string `json:"votes"` // Mapeia os ids dos usuarios e o valor é a opçao votada
	Connections []*websocket.Conn `json:"-"`     // Clientes conectados nesta enquete
}

type WSContract struct {
	MesaggeType         int               `json:"mensagge_type"`
	Vote                string            `json:"vote"`
	ChangeOptionsParams map[string]string `json:"change_options_params"`
}

type WSReponseContract struct {
	Poll        *Poll `json:"poll"`
	MesaggeType int   `json:"mesagge_type"`
	StatusCode  int   `json:"status_code"`
}

/*
message type:
1 = start connection
2 = vote
3 = change options
*/

type PostPollContract struct {
	Name     string   `json:"name"`
	Question string   `json:"question"`
	Options  []string `json:"options"`
}

type Claims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

// Variável global para armazenar as polls
var Pools = make(map[string]*Poll)

var JwtSecret = []byte(os.Getenv("JWT_SECRET"))

var WsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
