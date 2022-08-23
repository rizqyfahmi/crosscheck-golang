package authentication_persistent

import (
	"crosscheck-golang/app/features/authentication/data/model"

	"github.com/jmoiron/sqlx"
)

type AuthPersistent interface {
	Insert(userModel *model.UserModel) error
	GetByUsername(username *string) (*model.UserModel, error)
}

type AuthPersistentImpl struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) AuthPersistent {
	return &AuthPersistentImpl{
		db,
	}
}

func (s *AuthPersistentImpl) Insert(userModel *model.UserModel) error {
	_, err := s.db.NamedExec("INSERT INTO users (id, name, email, password, created_at, updated_at) VALUES (:id, :name, :email, :password, :created_at, :updated_at)", userModel)
	return err
}

func (s *AuthPersistentImpl) GetByUsername(username *string) (*model.UserModel, error) {
	userModel := model.UserModel{}

	row := s.db.QueryRowx("SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1", &username)
	err := row.StructScan(&userModel)

	if err != nil {
		return nil, err
	}

	return &userModel, nil
}
