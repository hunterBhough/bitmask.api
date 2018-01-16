package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/hunterBhough/go-doge/src/restapi/operations"
	"github.com/hunterBhough/go-doge/src/models"
	"github.com/hunterBhough/go-doge/src/cryptography"
)

func GetRecordsHandlerFunc(params operations.GetRecordsParams) middleware.Responder {

	return handle()
}

// loop over transmission_types
func buildRecords() []*models.Record {
	records := []*models.Record{}
	transactionTempPointer := models.Transaction{
		Amount:  1.12424912,
		Time: int64(12424912),
		Hash: "112424912",
		BlockHash: "112424912",
	}
	for i := 0; i < 2; i++ {
		record := cryptography.Decrypt(transactionTempPointer)
		records = append(records, record)
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

	for j := 0; j < 2; j++ {
		name := "NNDSS Food Borne"
		transmissionTypeRecords = buildTransmissionTypeRecords(buildRecords())

		transmissionType := models.TransmissionType{
			Name: name,
			Records: transmissionTypeRecords,
		}
		payload = append(payload, transmissionType)
	}


	return operations.NewGetRecordsOK().WithPayload(payload)
}