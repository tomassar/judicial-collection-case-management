package users

type SignUpReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
