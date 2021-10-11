package main

import (
	"github.com/Rohan12152001/CovidTrackerWebService/endpoints/covidtracker"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	covidData := covidtracker.New()
	covidData.SetRoutes(server)

	server.Run()
}