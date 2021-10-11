package db

import (
	"context"
	"fmt"
	"github.com/Rohan12152001/CovidTrackerWebService/managers/covidtracker/data"
	"github.com/Rohan12152001/CovidTrackerWebService/utils"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

type manager struct {
	db *sqlx.DB
}

func New() CovidDBManager {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return manager{
		db: db,
	}
}

func (m manager) GetAllData() ([]data.CovidDataStruct, error) {
	query := "Select date, cases, deaths, active, recovered from covidData"
	covidData := []data.CovidDataStruct{}

	err := m.db.Select(&covidData, query)
	if err != nil {
		return []data.CovidDataStruct{}, err
	}

	return covidData, nil
}

func (m manager) ReadData(ctx context.Context, DateFromParam string) (*data.CovidDataStruct, error) {
	query := "Select date, cases, deaths, active, recovered from covidData where date=$1"
	datas := []*data.CovidDataStruct{}

	err := m.db.Select(&datas, query, DateFromParam)
	if err != nil {
		return nil, err
	}

	if len(datas) == 0 {
		return nil, utils.NoRowsFound
	}

	return datas[0], nil
}

func (m manager) CreateData(dataPayload data.CovidDataStruct) (dataId int, err error) {
	query := "Insert into covidData (date, cases, deaths, active, recovered) values($1, $2, $3, $4, $5) RETURNING id;"

	Id := -1
	err = m.db.QueryRow(query, dataPayload.Date, dataPayload.Cases, dataPayload.Deaths,dataPayload.Active, dataPayload.Recovered).Scan(&Id)
	if err != nil {
		if err, ok := err.(*pq.Error); ok {
			if err.Code == "23505" {
				//handle duplicate insert
				return Id, utils.RowAlreadyExist
			}
		}
		return Id, err
	}

	return Id, nil
}

func (m manager) UpdateData(ctx context.Context, dataPayload data.CovidDataStruct, DateFromParam string) (err error) {
	// Begin transaction
	tx, err := m.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}

	// Then create
	creatQuery := "UPDATE covidData SET cases=$1, deaths=$2, active=$3, recovered=$4 where date=$5;"

	_, err = tx.Exec(creatQuery, dataPayload.Cases, dataPayload.Deaths, dataPayload.Active, dataPayload.Recovered, DateFromParam)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (m manager) DeleteData(ctx context.Context, DateFromParam string) (err error) {
	// Begin transaction
	tx, err := m.db.BeginTxx(context.Background(), nil)
	if err != nil {
		return err
	}

	// Then create
	creatQuery := "DELETE FROM covidData where date=$1;"

	_, err = tx.Exec(creatQuery, DateFromParam)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}



