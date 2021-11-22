package db

import (
	"context"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/covidtracker/data"
)


type CovidDBManager interface {
	GetAllData() ([]data.CovidDataStruct, error)
	ReadData(ctx context.Context, DateFromParam string) (*data.CovidDataStruct, error)
	CreateData(dataPayload data.CovidDataStruct)(dataId int, err error)
	UpdateData(ctx context.Context, dataPayload data.CovidDataStruct, DateFromParam string) (err error)
	DeleteData(ctx context.Context, DateFromParam string) (err error)
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "admin"
	dbname   = "CovidTracker"
)