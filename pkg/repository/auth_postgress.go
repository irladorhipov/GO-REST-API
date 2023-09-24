package repository

import  (
	"fmt"
	"application"
	"github.com/jmoiron/sqlx"
	)
type AuthPostgress struct {
	db *sqlx.DB
}

func NewAuthPostgress(db *sqlx.DB) *AuthPostgress {
	return &AuthPostgress{db:db}
}

func (r *AuthPostgress) CreateUser(user application.User) (int, error) {
	var id int
	
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id")
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}