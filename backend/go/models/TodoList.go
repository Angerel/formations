package models

type TodoList struct {
	Id   string      `json:"id"`
	List []*TodoItem `json:"list"`
}
