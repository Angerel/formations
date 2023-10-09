package models

import "github.com/gorilla/websocket"

type ReceivedMessage struct {
	Action  string          `json:"action"`
	Sender  *websocket.Conn `json:"-"`
	Options string          `json:"options"`
}
