package model



type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Id int `json:"id"`
}

type LoginPayload struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Id int `json:"id"`
}

type ErrResponse struct {
	Err error `json:"err"`
	Code int `json:"code"`
	StatusText string `json:"statusText"`
	ErrorText string `json:"errorText"`
}
