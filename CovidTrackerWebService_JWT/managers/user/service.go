package user

import (
	datas "github.com/Rohan12152001/CovidTrackerWebService/managers/user/data"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/user/db"
	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"time"
)

type manager struct {
	userDb db.UserDb
}

// secret key
var jwtkey = []byte("secret_key")

// Claims struct for payload
type Claims struct {
	UserId string `json:"user_id"`
	jwt.StandardClaims
}

func New() UserManager {
	return manager{
		userDb: db.New(),
	}
}

func (m manager) CreateUser(fname, email, password string) (string, error) {
	user := datas.User{
		Fname:       fname,
		Email:       email,
		Password:    password,
		Id: 		 uuid.New().String(),
	}
	err := m.userDb.CreateUser(user)
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

func (m manager) GetUserbyId(id string) (datas.User, error) {
	return m.userDb.GetUserbyID(id)
}

func (m manager) LoginUser(email, password string) (string, int, error) {
	// check email & password in DB
	userObject, err := m.userDb.GetUserForLogin(email, password)
	if err != nil {
		return "", 0, err
	}

	// form the jwt token
	expirationTime := 5 * 60 // 5 minutes time

	claims := &Claims{
		UserId: userObject.Id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", 0, err
	}

	// return tokenStr to set the cookie
	return tokenString, expirationTime, nil
}