package models

type Message struct {
	Id        string `json:"id"`
	Message   string `json:"message"`
	Timestamp string `json:"timestamp"`
	User      *User  `json:"user"`
}
