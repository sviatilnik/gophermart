package user

import (
	"github.com/google/uuid"
	"github.com/sviatilnik/gophermart/internal/domain/user/value_objects"
)

type User struct {
	ID       string
	Login    value_objects.Login
	Password value_objects.Password
}

func NewUser(login, password string) (*User, error) {

	log, err := value_objects.NewLogin(login)
	if err != nil {
		return nil, err
	}

	pass, err := value_objects.NewPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       uuid.NewString(),
		Login:    log,
		Password: pass,
	}, nil
}
