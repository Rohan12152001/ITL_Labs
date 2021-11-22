package covidtracker

import (
	"encoding/json"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/covidtracker"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/covidtracker/data"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/middleware"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
	"io/ioutil"
)

type CovidData struct {
	covidManager   covidtracker.CovidManager
	authManager    middleware.AuthManager
}

var logger = logrus.New()

func New() CovidData {
	return CovidData{
		covidManager: covidtracker.New(),
		authManager:  middleware.New(),
	}
}

func getParam(c *gin.Context, paramName string) string {
	return c.Params.ByName(paramName)
}

func (C CovidData) SetRoutes(router *gin.Engine) {
	router.GET("/allData", C.authManager.AuthMiddleWareWithUser, C.GetAllData)
	router.GET("/data/:date",C.authManager.AuthMiddleWareWithUser, C.ReadData)
	router.POST("/data", C.authManager.AuthMiddleWareWithUser, C.CreateData)
	router.PUT("/data/:date", C.authManager.AuthMiddleWareWithUser, C.UpdateData)
	router.DELETE("/data/:date", C.authManager.AuthMiddleWareWithUser, C.DeleteData)
}

// GetAllData Handler
func (C CovidData) GetAllData(context *gin.Context) {
	// Call manager
	items, err := C.covidManager.GetAllData(context)
	if err != nil {
		// errors
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"data": items,
	})
}

// ReadData Handler
func (C CovidData) ReadData(context *gin.Context) {
	dateFromParam := getParam(context, "date")

	// Call manager
	singleData, err := C.covidManager.ReadData(context, dateFromParam)
	if err != nil {
		if xerrors.Is(err, covidtracker.ItemNotFound) {
			context.JSON(404, gin.H{
				"item": "Not found",
			})
			return
		}
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"data": singleData,
	})
}

// CreateData Handler
func (C CovidData) CreateData(context *gin.Context) {
	// Unmarshal the payload
	dataPayload := data.CovidDataStruct{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &dataPayload)
	if err != nil {
		logger.Error(err)
		context.AbortWithStatus(500)
		return
	}

	// Manager
	dataId, err := C.covidManager.CreateData(context, dataPayload)
	if err != nil {
		if xerrors.Is(err, covidtracker.AlreadyExist) {
			// TODO: Very new status code
			context.JSON(409, gin.H{
				"dataId": "Data already exists",
			})
			context.AbortWithStatus(409)
		}
		context.AbortWithStatus(500)
		return
	}

	context.JSON(200, gin.H{
		"dataId": dataId,
	})

}

// UpdateData Handler
func (C CovidData) UpdateData(context *gin.Context) {
	// Unmarshal data
	dataPayload := data.CovidDataStruct{}
	b, err := ioutil.ReadAll(context.Request.Body)

	err = json.Unmarshal(b, &dataPayload)
	if err != nil {
		logger.Error("err: ", err)
		context.AbortWithStatus(500)
		return
	}

	// Get ID from pathParam
	DateFromParam := getParam(context, "date")

	// Manager
	_, err = C.covidManager.UpdateData(context, dataPayload, DateFromParam)
	if err != nil {
		if xerrors.Is(err, covidtracker.ItemNotFound) {
			context.JSON(404, gin.H{
				"data": "Not found",
			})
			return
		}
		context.AbortWithStatus(500)
		return
	}

	context.JSON(200, gin.H{
		"data": "Updated",
	})
}

// DeleteData Handler
func (C CovidData) DeleteData(context *gin.Context) {
	// Get ID from pathParam
	DateFromParam := getParam(context, "date")

	// Manager
	_, err := C.covidManager.DeleteData(context, DateFromParam)
	if err != nil {
		if xerrors.Is(err, covidtracker.ItemNotFound) {
			context.JSON(404, gin.H{
				"data": "Not found",
			})
			return
		}
		context.AbortWithStatus(500)
	}

	context.JSON(200, gin.H{
		"data": "Deleted",
	})
}