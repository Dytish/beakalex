package models

import (
	"time"
)

type BotBody struct {
	nodes interface{}
	edges interface{}
}

type Bot struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Id_user int     `json:"id_user" `
	Token   string  `json:"token" `
	Body    []byte `json:"body" `
}
type BotFront struct {
	ID        uint `json:"id" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`

	Id_user int    `json:"id_user" `
	Token   string `json:"token" `
	Body    interface{} `json:"body" `
}

type Bots struct {
	Id_user int `json:"id_user" `
	Bot     interface{}
}
