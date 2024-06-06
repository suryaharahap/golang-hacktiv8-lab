package models

import "time"

type User struct {
	ID             int
	Name           string
	Email          string
	PasswordHash   string
	Occupation     string
	AvatarFileName string
	Age            int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
