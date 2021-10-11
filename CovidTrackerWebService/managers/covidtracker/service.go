package covidtracker

import (
	"context"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/covidtracker/data"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/covidtracker/db"
	"github.com/Rohan12152001/CovidTrackerWebService/utils"
	"github.com/sirupsen/logrus"
	"golang.org/x/xerrors"
)

type manager struct {
	db db.CovidDBManager
}

func New() CovidManager {
	return manager{
		db: db.New(),
	}
}

var logger = logrus.New()

func (m manager) GetAllData(ctx context.Context) ([]data.CovidDataStruct, error) {
	// Call db
	datas, err := m.db.GetAllData()
	if err != nil {
		logger.Error("err: ", err)
		return nil, err
	}
	return datas, nil
}

func (m manager) ReadData(ctx context.Context, DateFromParam string) (*data.CovidDataStruct, error) {
	// Call db
	singleData, err := m.db.ReadData(ctx, DateFromParam)
	if err != nil {
		if xerrors.Is(err, utils.NoRowsFound) {
			logger.Error("err: ", err)
			return nil, ItemNotFound
		}
		logger.Error("err: ", err)
		return nil, err
	}
	return singleData, nil
}

func (m manager) CreateData(ctx context.Context, dataPayload data.CovidDataStruct) (dataId int, err error) {
	// Call db
	dataId, err = m.db.CreateData(dataPayload)
	if err != nil {
		if xerrors.Is(err, utils.RowAlreadyExist) {
			logger.Error("err: ", err)
			return -1, AlreadyExist
		}
		logger.Error("err: ", err)
		return -1, err
	}

	return dataId, nil
}

func (m manager) UpdateData(ctx context.Context, dataPayload data.CovidDataStruct, DateFromParam string) (ok bool, err error) {
	// Check if item exists
	_, err = m.db.ReadData(ctx, DateFromParam)
	if err != nil {
		if xerrors.Is(err, utils.NoRowsFound) {
			logger.Error("err: ", err)
			return false, ItemNotFound
		}
		logger.Error("err: ", err)
		return false, err
	}

	// pass final item to DB
	err = m.db.UpdateData(ctx, dataPayload, DateFromParam)
	if err != nil {
		logger.Error("err: ", err)
		return false, err
	}

	return true, nil
}

func (m manager) DeleteData(ctx context.Context, DateFromParam string) (ok bool, err error) {
	// Check if item exists
	_, err = m.db.ReadData(ctx, DateFromParam)
	if err != nil {
		if xerrors.Is(err, utils.NoRowsFound) {
			logger.Error("err: ", err)
			return false, ItemNotFound
		}
		logger.Error("err: ", err)
		return false, err
	}

	// pass DateParam to DB
	err = m.db.DeleteData(ctx, DateFromParam)
	if err != nil {
		logger.Error("err: ", err)
		return false, err
	}

	return true, nil
}

