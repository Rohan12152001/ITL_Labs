package user

import (
	datas "github.com/Rohan12152001/CovidTrackerWebService/managers/user/data"
)

type UserManager interface {
	CreateUser(fname, email, password string) (string, error)
	GetUserbyId(id string) (datas.User, error)
	LoginUser(email, password string) (string, int, error)
}