package destinationService

import (
	"encoding/json"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type DestinationRepository struct {
	dbw *sqlx.DB
	dbr *sqlx.DB
}

func InitRepository(dbw *sqlx.DB, dbr *sqlx.DB) *DestinationRepository {
	return &DestinationRepository{
		dbw: dbw,
		dbr: dbr,
	}
}

func (ur *DestinationRepository) CreateDestination(usr *DestinationCreateRequest) error {
	query := `INSERT INTO dbo.Destinations (name, description, photo, created_at) 
			  VALUES (:name, :description, :photo, getdate())`
	_, err := ur.dbw.NamedExec(query, usr)
	if err != nil {
		fmt.Println(err)
		return err
	}

	jsonBytes, err := json.Marshal(usr)
	if err != nil {
		panic(err)
	}
	logrus.Infof(`(CreateDestination) -  %v`, string(jsonBytes))
	logrus.Info("(CreateDestination) - Successfully Created a Destination")

	return nil
}

func (ur *DestinationRepository) FindUserByUsername(username string) (*Destination, error) {
	query := fmt.Sprintf(`select * from dbo.Users where username = '%v'`, username)
	var Users Destination
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
