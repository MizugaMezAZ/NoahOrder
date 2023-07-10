package repository

import (
	"github.com/jmoiron/sqlx"
)

type IAuthRepository interface {
}

type authRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) IAuthRepository {
	return &authRepository{db}
}

// ----------------------------------

func (a *authRepository) Login(account string) (bool, error) {
	var pwd string

	query := `SELECT 
					password 
				FROM 
					user
				WHERE
					account = ?`

	if err := a.db.Get(&pwd, query, account); err != nil {
		return false, err
	}

	return true, nil
}
