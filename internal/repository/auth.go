package repository

import (
	"goAuthService/models"
	"goAuthService/consts"

	"fmt"
	"strings"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) GetRoles(id int) ([]string, error) {
	var roles []string
	query := fmt.Sprintf("select r.role_name from %s r join %s c on c.role_id = r.id where c.user_id = $1", consts.RolesTable, consts.UsersRolesTable)

	rows, err := r.db.Query(query, id)
	if err != nil {
		return roles, err
	}

	for rows.Next() {
		var role string
		err := rows.Scan(&role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, err
}

func (r *AuthRepository) GetRoleId(role string, userId int) (int, error) {
	var id int
	table := "t_" + strings.ToLower(role)

	query := fmt.Sprintf("select id from %s where user_id = $1", table)
	err := r.db.Get(&id, query, userId)
	return id, err
}

func (r *AuthRepository) GetUser(login string) (models.User, error) {
	var user models.User
	queryParam := "username"
	if strings.Contains(login, "@") {
		queryParam = "email"
	}
	query := fmt.Sprintf("select id, username, password from %s where %s = $1", consts.UsersTable, queryParam)
	err := r.db.Get(&user, query, login)
	if err != nil {
		if strings.Contains(err.Error(), "no rows") {
			return models.User{}, fmt.Errorf("there is no such user with username/email: %s", login)
		}
		return models.User{}, err
	}
	return user, nil
}

func (r *AuthRepository) CreateUser(user models.User, tx *sqlx.Tx) (int, error) {
	var userId int
	query := fmt.Sprintf("insert into %s (username, password, name, surname, email, phone) values ($1, $2, $3, $4, $5, $6) returning id", consts.UsersTable)
	row := tx.QueryRow(query, user.Username, user.Password, user.Name, user.Surname, user.Email, user.Phone)
	if err := row.Scan(&userId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return userId, nil
}
