package auth_persistent_data_test

import (
	"database/sql"
	"errors"
	"log"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"crosscheck-golang/app/exception"
	"crosscheck-golang/app/features/authentication/data/model"
	authentication_persistent "crosscheck-golang/app/features/authentication/data/source"
)

var _ = Describe("AuthPersistentData", func() {

	var mockUserModel *model.UserModel
	var authPersistent authentication_persistent.AuthPersistent
	var mockDB *sql.DB
	var mock sqlmock.Sqlmock
	var sqlxDB *sqlx.DB

	BeforeEach(func() {
		var err error
		mockDB, mock, err = sqlmock.New()

		if err != nil {
			log.Println("Error sqlmock!")
			log.Fatal(err)
		}

		sqlxDB = sqlx.NewDb(mockDB, "sqlmock")

		authPersistent = authentication_persistent.New(sqlxDB)

		mockUserModel = &model.UserModel{
			Id:        time.Now().Format("20060102150405"),
			Name:      "rizqyfahmi",
			Email:     "rizqyfahmi@email.com",
			Password:  "HelloPassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
	})

	Describe("Insert", func() {
		Context("When insert the registration data into database", func() {
			It("returns error", func() {
				defer mockDB.Close()

				mock.ExpectExec("INSERT INTO users").WillReturnError(errors.New(exception.ErrorDatabase))
				err := authPersistent.Insert(mockUserModel)

				Expect(err).Should(HaveOccurred())
				Expect(err).Should(MatchError(exception.ErrorDatabase))
			})

			It("returns nil", func() {
				defer mockDB.Close()

				mock.ExpectExec("INSERT INTO users (.+)").WithArgs(mockUserModel.Id, mockUserModel.Name, mockUserModel.Email, mockUserModel.Password, mockUserModel.CreatedAt, mockUserModel.UpdatedAt).WillReturnResult(sqlmock.NewResult(1, 1))
				err := authPersistent.Insert(mockUserModel)

				Expect(err).Should(Succeed())
			})
		})
	})
})
