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
func buildRecords(wallet models.Wallet) []*models.Record {
	records := []*models.Record{}
	for i := 0; i < wallet.TotalTxs; i++ {
		record := cryptography.Decrypt(wallet.Txs[i].Incoming.Value)
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
	//input
	wallets := DogeHandler()
	//output
	transmissionTypeRecords := models.TransmissionTypeRecords{}

	for i := 0; i < len(wallets); i++ {
		if wallets[i].TotalTxs != 0 {
			transmissionTypeRecords = buildTransmissionTypeRecords(buildRecords(wallets[i]))
		}

		transmissionType := models.TransmissionType{
			Name:    wallets[i].Name,
			Records: transmissionTypeRecords,
		}
		payload = append(payload, transmissionType)
	}


	return operations.NewGetRecordsOK().WithPayload(payload)
}