package controllers

import "golang-hactiv8-lab/final-project/mygram/models"

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"string"`
	Age        int    `json:"age"`
	Token      string `json:"token"`
}

func FormatUser(user models.User, token string) UserFormatter {
	formatter := UserFormatter{
		ID:         user.ID,
		Name:       user.Name,
		Occupation: user.Occupation,
		Age:        user.Age,
		Token:      token,
	}

	return formatter
}
