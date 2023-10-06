package models

import "github.com/gorilla/websocket"

type ConnectedUser struct {
	User       User            `json:"user"`
	Connection *websocket.Conn `json:"connection"`
}
