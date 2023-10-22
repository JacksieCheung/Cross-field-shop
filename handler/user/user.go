package user

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type EmailValidateReq struct {
	Email string `json:"email"`
}

type RegisterReq struct {
	Email        string `json:"email"`
	Password     string `json:"password"`
	Name         string `json:"name"`
	ValidateCode string `json:"validate_code"`
}
