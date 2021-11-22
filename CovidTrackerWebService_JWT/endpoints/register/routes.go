package register

import (
	"encoding/json"
	"fmt"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/user"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type Register struct {
	usermanager user.UserManager
}

type UserRegisterPayload struct {
	Fname       string `json:"fname"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

func New() Register {
	return Register{
		usermanager: user.New(),
	}
}

func (R *Register) SetRoutes(engine *gin.Engine) {
	engine.POST("/register", R.RegisterCall)
}

// RegisterCall for endpoint layer
func (R *Register) RegisterCall(context *gin.Context) {
	var payload UserRegisterPayload

	bytesRecieved, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		context.Error(err)
		return
	}

	err = json.Unmarshal(bytesRecieved, &payload)
	if err != nil {
		context.Error(err)
		return
	}

	UserId, err := R.usermanager.CreateUser(payload.Fname,
		payload.Email,
		payload.Password)
	if err != nil {
		context.Error(err)
		context.Status(500)
		return
	}

	fmt.Println("User created !")

	context.JSON(200, gin.H{
		"ID": UserId,
	})
}