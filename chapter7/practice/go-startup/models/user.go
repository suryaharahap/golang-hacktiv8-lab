package models

import "time"

type User struct {
	ID             int
	Name           string
	Occupation     string
	Email          string
	PasswordHash   string
	AvatarFileName string
	Role           string
	CratedAt       time.Time
	UpdatedAt      time.Time
}
