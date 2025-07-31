package user

import (
	"context"
	"database/sql"
	"github.com/sviatilnik/gophermart/internal/domain/user"
	"github.com/sviatilnik/gophermart/internal/domain/user/value_objects"
)

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *user.User) error {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresUserRepository) FindByLogin(ctx context.Context, login value_objects.Login) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresUserRepository) Exists(ctx context.Context, login value_objects.Login) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresUserRepository) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}
