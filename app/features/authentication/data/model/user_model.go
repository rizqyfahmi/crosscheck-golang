package model

import "time"

type UserModel struct {
	Id        string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}
