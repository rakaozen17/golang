package userService

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	dbw *sqlx.DB
	dbr *sqlx.DB
}

func InitRepository(dbw *sqlx.DB, dbr *sqlx.DB) *UserRepository {
	return &UserRepository{
		dbw: dbw,
		dbr: dbr,
	}
}

func (ur *UserRepository) RegisterAccount(usr UserModel) error {
	query := `INSERT INTO Users (username, email, password, first_name, last_name, address, phone_number, created_at, is_active) 
			  VALUES (:username, :email, :password, :first_name, :last_name, :address, :phone_number, getdate(), 1)`
	_, err := ur.dbw.NamedExec(query, usr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonBytes, err := json.Marshal(usr)
	if err != nil {
		panic(err)
	}
	logrus.Infof(`(RegisterAccount) -  %v`, string(jsonBytes))
	logrus.Info("(RegisterAccount) - Successfully registered an Account")

	return nil
}

func (ur *UserRepository) FindUserByUsername(username string) (*UserModel, error) {
	query := fmt.Sprintf(`select * from dbo.Users where username = '%v'`, username)
	var Users UserModel
	err := ur.dbw.Get(&Users, query)
	if err != nil {
		fmt.Println(err)
		return &Users, err
	}

	jsonBytes, err := json.Marshal(Users)
	if err != nil {
		panic(err)
	}
	logrus.Infof(`(Login) -  %v`, string(jsonBytes))
	logrus.Info("(Login) - Successfully get the users")

	return &Users, nil
}
