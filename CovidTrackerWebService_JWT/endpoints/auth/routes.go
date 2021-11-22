package auth

import (
	"encoding/json"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/user"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Authendpoints struct {
	usermanager user.UserManager
}

type UserLoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func New() Authendpoints {
	return Authendpoints{
		usermanager: user.New(),
	}
}

func (A Authendpoints) SetRoutes(router *gin.Engine) {
	router.POST("/login", A.LoginHandler)
}

// LoginHandler for endpoint layer
func (A Authendpoints) LoginHandler(context *gin.Context) {
	var LoginPayload UserLoginPayload

	// b is bytes
	b, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.Error(err)
		context.Status(500)
		return
	}

	err = json.Unmarshal(b, &LoginPayload)
	if err != nil {
		context.Error(err)
		context.Status(500)
		return
	}

	// jwt begin
	tokenString, expirationTime, err := A.usermanager.LoginUser(LoginPayload.Email, LoginPayload.Password)
	if err != nil {
		context.Error(err)
		context.Status(401) // what should we return status code (since it can be 401 or 500) ?
		return
	}

	context.SetCookie(
		"token",
		tokenString,
		expirationTime,
		"/",
		"localhost",
		false,
		false,
	)
}