package user

import (
	"context"
	"errors"
	"github.com/sviatilnik/gophermart/internal/domain/user/value_objects"
)

var (
	ErrUserNotFound      = errors.New("user not found")
	ErrUserAlreadyExists = errors.New("user already exists")
)

type Repository interface {
	Save(ctx context.Context, user *User) error
	FindByID(ctx context.Context, id string) (*User, error)
	FindByLogin(ctx context.Context, login value_objects.Login) (*User, error)
	Exists(ctx context.Context, login value_objects.Login) (bool, error)
	Delete(id string) error
}
