package repository

import (
	"errors"
	"fmt"
	"goAuthService/consts"
	"goAuthService/models"
	"strings"

	"github.com/rs/zerolog/log"
)

func (r *AuthRepository) CreateClient(client models.User) (int, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		log.Error().Msg(err.Error())
		return 0, err
	}
	var clientId int
	userId, err := r.CreateUser(client, tx)

	if err != nil {
		log.Error().Msg(err.Error())
		msg := err.Error()
		if strings.Contains(msg, "duplicate") {
			if strings.Contains(msg, "email") {
				return -1, errors.New("there is already exist account with such email")
			}
			if strings.Contains(msg, "username") {
				return -1, errors.New("there is already exist account with such username")
			}
		}
		tx.Rollback()
		return 0, err
	}

	query := fmt.Sprintf("insert into %s (user_id) values($1) returning id", consts.ClientsTable)
	row := tx.QueryRowx(query, userId)

	if err := row.Scan(&clientId); err != nil {
		log.Error().Msg(err.Error())
		tx.Rollback()
		return 0, err
	}

	query = fmt.Sprintf("insert into %s (user_id, role_id) values ($1, $2)", consts.UsersRolesTable)
	log.Info().Msg(query)
	_, err = tx.Exec(query, userId, consts.ClientRoleId)

	if err != nil {
		tx.Rollback()
		log.Error().Msg(err.Error())
		return 0, err
	}

	return clientId, tx.Commit()
}