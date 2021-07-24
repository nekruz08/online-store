package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nekruz08/online-store/models"
	"github.com/jmoiron/sqlx"
)


type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

//CreateUser - создает пользователя и хеширует пароль
func (r AuthPostgres) CreateUser(user models.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s(name, username, password_hash,role_id) values($1, $2, $3, 2) RETURNING id", userTable)
	err := r.db.QueryRow(query, user.Name, user.Username, user.Password).Scan(&id)
	if err != nil {
		return 0, err

	}

	return id, nil

}

// GetUser - возвращает юзера по логину и паролю
func (r *AuthPostgres) GetUser(username, password string) (models.User, error) {
	var user models.User

	query := fmt.Sprintf("SELECT id FROM %s WHERE username=$1 AND password_hash=$2", userTable)
	err := r.db.Get(&user, query, username, password)

	return user, err
}
//AdminChecker  - проверка на Роль Админа
func (r *AuthPostgres) AdminChecker(userId int) (bool, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE id=$1 AND role_id=1", userTable)

	err := r.db.Get(&user, query, userId)
	if err == sql.ErrNoRows {
		return false, errors.New("у вас нет доступа, чтобы добавит  товар")
	}
	if err != nil {
		return false, err
	}

	return true, err

}
