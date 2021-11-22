package covidtracker

import (
	"context"
	"fmt"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/covidtracker/data"
)

var (
	ItemNotFound = fmt.Errorf("item not found")
	AlreadyExist = fmt.Errorf("row already exists")
	)

type CovidManager interface {
	GetAllData(ctx context.Context) ([]data.CovidDataStruct, error)
	ReadData(ctx context.Context, DateFromParam string) (*data.CovidDataStruct, error)
	CreateData(ctx context.Context, dataPayload data.CovidDataStruct) (dataId int, err error)
	UpdateData(ctx context.Context, dataPayload data.CovidDataStruct, DateFromParam string) (ok bool, err error)
	DeleteData(ctx context.Context, DateFromParam string) (ok bool, err error)
}
