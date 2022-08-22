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
	authpersistent "crosscheck-golang/app/features/authentication/data/source/persistent"
)

var _ = Describe("AuthPersistentData", func() {

	var mockUserModel *model.UserModel
	var mockUsername string
	var authPersistent authpersistent.AuthPersistent
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

		authPersistent = authpersistent.New(sqlxDB)

		mockUserModel = &model.UserModel{
			Id:        time.Now().Format("20060102150405"),
			Name:      "rizqyfahmi",
			Email:     "rizqyfahmi@email.com",
			Password:  "HelloPassword",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		mockUsername = "rizqyfahmi@email.com"
	})

	Context("Insert", func() {
		Describe("UserModel as the parameter", func() {
			When("Executing insert query", func() {
				It("returns an error", func() {
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

	Context("Get user model by username", func() {
		Describe("Username as the parameter", func() {
			When("Executing select query where username is equal to the parameter", func() {
				It("returns an error", func() {
					defer mockDB.Close()
					mock.ExpectQuery("SELECT id, name, email, password, created_at, updated_at FROM users WHERE users.id = (.+)").
						WithArgs(mockUsername).
						WillReturnError(errors.New(exception.ErrorDatabase))

					result, err := authPersistent.GetByUsername(&mockUsername)

					Expect(err).Should(HaveOccurred())
					Expect(err).Should(MatchError(exception.ErrorDatabase))
					Expect(result).Should(BeNil())
				})

				It("returns an user model", func() {
					defer mockDB.Close()

					mockQueryResult := sqlmock.NewRows([]string{"id", "name", "email", "password", "created_at", "updated_at"}).
						AddRow(mockUserModel.Id, mockUserModel.Name, mockUserModel.Email, mockUserModel.Password, mockUserModel.CreatedAt, mockUserModel.UpdatedAt)

					mock.ExpectQuery("SELECT id, name, email, password, created_at, updated_at FROM users WHERE users.id = (.+)").
						WithArgs(mockUsername).
						WillReturnRows(mockQueryResult)

					result, err := authPersistent.GetByUsername(&mockUsername)

					Expect(err).Should(Succeed())
					Expect(result.Id).Should(Equal(mockUserModel.Id))
				})
			})
		})
	})
})
