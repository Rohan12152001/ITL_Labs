package db

import (
	"fmt"
	datas "github.com/Rohan12152001/CovidTrackerWebService/managers/user/data"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type dbClient struct {
	db *sqlx.DB
}

func New() UserDb {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return dbClient{
		db: db,
	}
}

func (d dbClient) CreateUser(user datas.User) error {
	query := "Insert into users (fname, email, password, id) values (:fname, :email, :password, :id)"
	_, err := d.db.NamedExec(query, &user)
	if err != nil {
		return err
	}

	fmt.Println("INSERT DONE!")
	return nil
}

func (d dbClient) GetUserbyID(Id string) (datas.User, error) {
	query := "Select * from users where id=$1"
	users := []datas.User{}
	err := d.db.Select(&users, query, Id)
	if err != nil {
		return datas.User{}, err
	}
	return users[0], nil
}

func (d dbClient) GetUserForLogin(email, password string) (datas.User, error) {
	query := "Select * from users where email=$1 and password=$2"
	users := []datas.User{}
	err := d.db.Select(&users, query, email, password)
	if err != nil {
		return datas.User{}, err
	}
	return users[0], nil
}