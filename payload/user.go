package payload

type User struct {
	iduser   int    `json:"iduser"`
	Username string `json:"username"`
	Password string `json:"password"`
}
