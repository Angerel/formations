package models

type FirstMessage struct {
	History []*Message `json:"history"`
	UserId  string     `json:"userId"`
}
