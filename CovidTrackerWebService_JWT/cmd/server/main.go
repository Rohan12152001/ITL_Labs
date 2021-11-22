package main

import (
	"github.com/Rohan12152001/CovidTrackerWebService/endpoints/auth"
	"github.com/Rohan12152001/CovidTrackerWebService/endpoints/covidtracker"
	"github.com/Rohan12152001/CovidTrackerWebService/endpoints/register"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	covidData := covidtracker.New()
	covidData.SetRoutes(server)

	authorise := auth.New()
	authorise.SetRoutes(server)

	reg := register.New()
	reg.SetRoutes(server)

	server.Run()
}