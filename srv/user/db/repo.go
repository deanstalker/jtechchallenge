package db

import (
	"database/sql"
	"errors"
	"strings"

	user "github.com/deanstalker/jtechchallenge/srv/user/proto"
)

type UserRepo interface {
	BatchCreate(users []*user.UserItem) (int64, error)
	Create(item *user.UserItem) (int64, error)
	Delete(username string) error
	ByUsername(username string) (*user.UserItem, error)
	Update(user *user.UserItem) error
}

type DefaultUserRepo struct {
	db *sql.DB
}

var _ UserRepo = (*DefaultUserRepo)(nil)

func NewUserRepo(db *sql.DB) *DefaultUserRepo {
	return &DefaultUserRepo{
		db: db,
	}
}

func (r *DefaultUserRepo) BatchCreate(users []*user.UserItem) (int64, error) {
	qs := QueryBatchAdd
	valuesMask := `(?, ?, ?, ?, ?, ?, ?)`
	valueParams := make([]string, len(users))

	for range users {
		valueParams = append(valueParams, valuesMask)
	}

	qs = qs + strings.Join(valueParams, ", ")

	stmt, err := r.db.Prepare(qs)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	params := make([]interface{}, len(users)*7)
	for _, u := range users {
		paramSet := make([]interface{}, 7)
		paramSet[0] = u.Username
		paramSet[1] = u.FirstName
		paramSet[2] = u.LastName
		paramSet[3] = u.Email
		paramSet[4] = u.Password
		paramSet[5] = u.Phone
		paramSet[6] = u.UserStatus
		params = append(params, paramSet...)
	}

	result, err := stmt.Exec(params...)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *DefaultUserRepo) Create(item *user.UserItem) (int64, error) {
	stmt, err := r.db.Prepare(QueryAdd)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		item.Username,
		item.FirstName,
		item.LastName,
		item.Email,
		item.Password,
		item.Phone,
		item.UserStatus,
	)
	if err != nil {
		return 0, err
	}

	return result.RowsAffected()
}

func (r *DefaultUserRepo) Delete(username string) error {
	stmt, err := r.db.Prepare(QueryDelete)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}

	return nil
}

func (r *DefaultUserRepo) ByUsername(username string) (*user.UserItem, error) {
	stmt, err := r.db.Prepare(QueryGetByUsername)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)

	var token sql.NullString

	u := &user.UserItem{}
	if err = row.Scan(
		&u.Id,
		&u.Username,
		&u.FirstName,
		&u.LastName,
		&u.Email,
		&u.Password,
		&u.Phone,
		&u.UserStatus,
		&token,
	); err != nil {
		return nil, err
	}

	if token.Valid {
		u.Token = token.String
	}

	return u, nil
}

func (r *DefaultUserRepo) Update(user *user.UserItem) error {
	stmt, err := r.db.Prepare(QueryUpdate)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Username, user.FirstName, user.LastName, user.Email, user.Password, user.Phone, user.UserStatus, user.Id)
	if err != nil {
		return err
	}

	total, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if total == 0 {
		return errors.New("record was not updated")
	}

	return nil
}
