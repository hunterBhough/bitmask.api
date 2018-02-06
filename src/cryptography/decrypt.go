package cryptography

/* Takes in the value from the transaction
	Validates Input
	Returns Record
*/
// Dependency: response-handler.go

import (
	"github.com/hunterBhough/go-doge/src/models"
	"strconv"
)

func Decrypt(value string, id int64) (*models.Record, models.Error) {
	var pregnancy,
	hospitalized,
	death,
	innc,
	outbreak string

	var duration, age int64

	var err models.Error
	var record models.Record

	pregnancy, err = getPregnancy(string(value[2]))
	if err.Status != 0 {
		return &record, err
	}

	hospitalized, err = getBinaryAnswer(string(value[3]))
	if err.Status != 0 {
		return &record, err
	}

	duration, err = getDuration(string(value[4]))
	if err.Status != 0 {
		return &record, err
	}

	death, err = getBinaryAnswer(string(value[5]))
	if err.Status != 0 {
		return &record, err
	}

	age, err = getAge(string(value[6]), string(value[7]))
	if err.Status != 0 {
		return &record, err
	}

	innc, err = getBinaryAnswer(string(value[8]))
	if err.Status != 0 {
		return &record, err
	}

	outbreak, err = getBinaryAnswer(string(value[9]))
	if err.Status != 0 {
		return &record, err
	}

	if err.Message == "" {
		record = models.Record{
			Age:                                  age,
			CaseOutbreakIndicator:                outbreak,
			Death:                                death,
			Duration:                             duration,
			Hospitalized:                         hospitalized,
			ID:                                   id,
			ImmediateNationalNotifiableCondition: innc,
			PregnancyStatus:                      pregnancy,
		}
	}

	return &record, err
}

func getBinaryAnswer(digit string) (string, models.Error) {
	var result string
	var err models.Error

	if digit == "1" {
		result = "No"
	} else if digit == "2" {
		result = "Yes"
	} else {
		err.Message += "invalid input \n"
		err.Status = 1
	}

	return result, err
}

func getPregnancy(digit string) (string, models.Error) {
	var pregnancy string
	var err models.Error

	if digit == "0" {
		pregnancy = "No"
	} else if digit == "1" {
		pregnancy = "Yes"
	} else {
		err.Message += "invalid pregnancy input \n"
		err.Status = 1
	}

	return pregnancy, err
}

func getDuration(digit string) (int64, models.Error) {
	var duration int64
	var err models.Error

	var durInt, durErr = strconv.Atoi(digit)
	if durErr != nil {
		err.Message += "invalid duration input \n"
		err.Status = 1
	} else {
		duration = int64(durInt)
	}

	return duration, err
}

func getAge(tens string, singles string) (int64, models.Error) {
	var age int64
	var err models.Error

	var ageInt, ageErr = strconv.Atoi(tens + singles)
	if ageErr != nil {
		err.Message += "invalid age input \n"
		err.Status = 1
	} else {
		age = int64(ageInt)
	}

	return age, err
}
