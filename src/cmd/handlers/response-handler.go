package handlers

import (
	"github.com/hunterBhough/go-doge/src/models"
	"github.com/hunterBhough/go-doge/src/cryptography"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"encoding/json"
	λ "github.com/hunterBhough/go-doge/src/errors"

)


func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
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
	body, err := json.Marshal(payload)
	if err != nil {
		λ.MarshalError(err)
	}
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
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

// swagger server code

//func GetRecordsHandlerFunc() middleware.Responder {
//	return handle()
//}

// loop over wallets and build records into each of them
//func handle() middleware.Responder {
//	payload := []models.TransmissionType{}
//	//input
//	wallets := DogeHandler()
//	//output
//	transmissionTypeRecords := models.TransmissionTypeRecords{}
//
//	for i := 0; i < len(wallets); i++ {
//		if wallets[i].TotalTxs != 0 {
//			transmissionTypeRecords = buildTransmissionTypeRecords(buildRecords(wallets[i]))
//		}
//
//		transmissionType := models.TransmissionType{
//			Name:    wallets[i].Name,
//			Records: transmissionTypeRecords,
//		}
//		payload = append(payload, transmissionType)
//	}
//
//	return operations.NewGetRecordsOK().WithPayload(payload)
//}