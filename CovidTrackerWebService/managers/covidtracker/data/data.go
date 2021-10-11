package data

type CovidDataStruct struct {
	Date string `json:"date"`
	Cases int `json:"cases"`
	Active int `json:"active"`
	Deaths int `json:"deaths"`
	Recovered int `json:"recovered"`
}