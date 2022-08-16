package authentication_persistent

import (
	"crosscheck-golang/app/features/authentication/data/model"
)

type AuthPersistent interface {
	Insert(userModel *model.UserModel) error
}
