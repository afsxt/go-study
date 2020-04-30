package auth_service

import "base-server/models"

//-----------------------------------------------------------------------------

type Auth struct {
	Username string
	Password string
}

func NewAuth(username, password string) *Auth {
	a := Auth{
		Username: username,
		Password: password,
	}

	return &a
}

func (a *Auth) Check() (bool, error) {
	return models.CheckAuth(a.Username, a.Password)
}
