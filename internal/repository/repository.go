package repository

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/lallison21/auth_service/internal/models"
)

type Repository struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Register(ctx context.Context, newUser *models.UserDao) (int, error) {
	query := `INSERT INTO users(username, password, email) VALUES ($1, $2, $3) RETURNING id`

	var userId int
	err := r.db.QueryRow(ctx, query, newUser.Username, newUser.Password, newUser.Email).Scan(&userId)
	if err != nil {
		return -1, err
	}

	return userId, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*models.UserDao, error) {
	queryBuilder := squirrel.Select("id", "username", "password", "email").
		From("users").
		Where(squirrel.Eq{"email": email})

	query, args, err := queryBuilder.PlaceholderFormat(squirrel.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	var user models.UserDao
	err = r.db.QueryRow(ctx, query, args...).
		Scan(&user.Id, &user.Username, &user.Password, &user.Email)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
