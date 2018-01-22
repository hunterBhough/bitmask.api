package cryptography

import (
	"github.com/hunterBhough/go-doge/src/models"
	"strconv"
)

/*	input
	------
	Transaction

	==========================================
	output
	-------
	Record
*/


func Decrypt(value string) *models.Record {
	var id = int64(1)
	var pregnancy string
	var hospitalized string
	var duration, durErr = strconv.Atoi(string(value[4]))
	if durErr != nil {
		duration = 0
	}
	var death string
	var age, ageErr = strconv.Atoi(string(value[6]) + string(value[7]))
	if ageErr != nil {
		age = 0
	}
	var innc string
	var outbreak string

	if string(value[2]) == "1"{
		pregnancy = "No"
	} else if string(value[2]) == "2" {
		pregnancy = "Yes"
	}
	if string(value[3]) == "1" {
		hospitalized = "No"
	} else if string(value[3]) == "2" {
		hospitalized = "Yes"
	}
	if string(value[5]) == "1" {
		death = "No"
	} else if string(value[5]) == "2" {
		death = "Yes"
	}
	if string(value[8]) == "1" {
		innc = "No"
	} else if string(value[8]) == "2" {
		innc = "Yes"
	}
	if string(value[9]) == "1" {
		outbreak = "No"
	} else if string(value[9]) == "2" {
		outbreak = "Yes"
	}


	record := models.Record{
		Age: int64(age),
		CaseOutbreakIndicator: outbreak,
		Death: death,
		Duration: int64(duration),
		Hospitalized: hospitalized,
		ID: id,
		ImmediateNationalNotifiableCondition: innc,
		PregnancyStatus: pregnancy,
	}

	return &record
}