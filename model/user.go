package model

type User struct {
	iduser   int    `json:"iduser"`
	Username string `json:"username"`
	Password string `json:"password"`
	Gender   int    `json:"gender"`
	Email    string `json:"email"`
}
