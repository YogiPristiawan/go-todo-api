package repositories

import (
	"context"
	"time"

	"go_todo_api/src/account/models"
	"go_todo_api/src/shared/databases"
)

// AccountRepository is an abstract contains methods
// required to provides account data
type AccountRepository interface {
	GetByUsername(username string) (account models.Account, err error)
	Create(account *models.Account) error

	GetProfileByUserId(userId int64) (profile models.Profile, err error)
}

// accountRepository provides methods for providing
// account data
type accountRepository struct {
	db databases.DB
}

// NewAccountRepository creates an instance of accountRepository
func NewAccountRepository() AccountRepository {
	return &accountRepository{}
}

// GetByUsername provide user account data
// by the given username
func (a *accountRepository) GetByUsername(username string) (account models.Account, err error) {
	var sql = `
		SELECT
			id, username, password, gender, birth_date
		FROM
			users
		WHERE
			username = $1`

	err = a.db.QueryRow(context.Background(), sql, username).Scan(
		&account.Id, &account.Username, &account.Password, &account.Gender, &account.BirthDate,
	)
	return
}

// Create handle an action to create user's account data
func (a *accountRepository) Create(account *models.Account) error {
	var sql = `
	INSERT INTO users
	(
		username, password, gender, birth_date, created_at, updated_at
	)
	VALUES
	(
		$1, $2, $3 $4, $5, $6
	) RETURNING id, username, password, gender, birth_date, created_at, updated_at`

	// create timestamps
	timestamps := time.Now().Unix()
	account.CreatedAt = timestamps
	account.UpdatedAt = timestamps

	err := a.db.QueryRow(context.Background(), sql,
		&account.Username, &account.Password, &account.Gender, &account.BirthDate,
		&account.CreatedAt, &account.UpdatedAt,
	).Scan(
		&account.Id, &account.Username, &account.Password, &account.Gender, &account.BirthDate, &account.CreatedAt, &account.UpdatedAt)
	return err
}

// GetProfileByUserId handle an action to get user's profile
// with the given user id
func (a *accountRepository) GetProfileByUserId(userId int64) (profile models.Profile, err error) {
	var sql = `
		SELECT
			id, username, gender, birth_date
		FROM
			users
		WHERE id = $1`

	err = a.db.QueryRow(context.Background(), sql, userId).Scan(
		&profile.Id, &profile.Username, &profile.Gender, &profile.BirthDate)
	return
}
