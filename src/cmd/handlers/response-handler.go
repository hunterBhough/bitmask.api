package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/hunterBhough/go-doge/src/restapi/operations"
	"github.com/hunterBhough/go-doge/src/models"
)

func GetRecordsHandlerFunc(params operations.GetRecordsParams) middleware.Responder {

	return handle()
}

// loop over transmission_types

func buildRecords() []*models.Record {
	records := []*models.Record{}
	for i := 0; i < 10; i++ {
		age := int64(29)
		outbreak := "No"
		death := "No"
		duration := int64(4)
		hospitalized := "Yes"
		id := int64(i)
		innc := "No"
		pregnancy := "Yes"


		record := models.Record{
			Age: age,
			CaseOutbreakIndicator: outbreak,
			Death: death,
			Duration: duration,
			Hospitalized: hospitalized,
			ID: id,
			ImmediateNationalNotifiableCondition: innc,
			PregnancyStatus: pregnancy,
		}
		records = append(records, &record)
	}
	return records
}

func buildTransmissionTypeRecords(records []*models.Record) models.TransmissionTypeRecords {
	transmissionTypeRecords := records
	return transmissionTypeRecords
}

// loop over wallets and build records into each of them
func handle() middleware.Responder {
	payload := []models.TransmissionType{}
	transmissionTypeRecords := models.TransmissionTypeRecords{}

	for j := 0; j < 4; j++ {
		name := "NNDSS Food Borne"
		transmissionTypeRecords = buildTransmissionTypeRecords(buildRecords())

		transmissionType := models.TransmissionType{
			Name: &name,
			Records: transmissionTypeRecords,
		}
		payload = append(payload, transmissionType)
	}


	return operations.NewGetRecordsOK().WithPayload(payload)
}