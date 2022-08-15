package authentication_data_source

import (
	"crosscheck-golang/app/features/authentication/data/model"
)

type AuthPersistentDataSource interface {
	Insert(userModel *model.UserModel) error
}
