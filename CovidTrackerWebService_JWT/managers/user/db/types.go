package db

import (
	datas "github.com/Rohan12152001/CovidTrackerWebService/managers/user/data"
)

type UserDb interface {
	CreateUser(user datas.User) error
	GetUserbyID(Id string) (datas.User, error)
	GetUserForLogin(email, password string) (datas.User, error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "CovidTracker"
)