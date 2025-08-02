package user

import (
	"context"
	"database/sql"
	"github.com/Masterminds/squirrel"
	"github.com/sviatilnik/gophermart/internal/domain/user"
	"github.com/sviatilnik/gophermart/internal/domain/user/value_objects"
)

type PostgresUserRepository struct {
	db      *sql.DB
	builder squirrel.StatementBuilderType
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{
		db:      db,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *user.User) error {
	query, _, err := r.builder.
		Insert("users").
		Columns("id", "login", "password").
		Values("?", "?", "?").
		ToSql()

	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, user.ID, user.Login, user.Password)
	if err != nil {
		return err
	}

	return nil
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	query, _, err := r.builder.
		Select("id", "login", "password").
		From("users").Where("id = ?").
		ToSql()

	if err != nil {
		return nil, err
	}

	usr := &user.User{}

	err = r.db.
		QueryRowContext(ctx, query, id).
		Scan(usr.ID, usr.Login, usr.Password)

	if err != nil {
		return nil, err
	}

	return usr, nil
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
