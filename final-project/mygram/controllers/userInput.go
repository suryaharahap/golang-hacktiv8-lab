package controllers

type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Age        int    `json:"age" binding:"required"`
}

type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}