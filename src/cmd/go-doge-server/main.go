package main

import (
	"github.com/hunterBhough/go-doge/src/models"
	"github.com/hunterBhough/go-doge/src/cmd/handlers"
	"github.com/hunterBhough/go-doge/src/cryptography"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/events"

	//"encoding/json"
	λ "github.com/hunterBhough/go-doge/src/errors"

	"fmt"
	//"github.com/go-openapi/runtime/middleware"
	//"github.com/hunterBhough/go-doge/src/restapi/operations"
	"encoding/json"
)

func Handler() (events.APIGatewayProxyResponse, error) {
	var payload []models.TransmissionType
	//input
	wallets, errors := handlers.DogeHandler()
	//output
	transmissionTypeRecords := models.TransmissionTypeRecords{}

	for i := 0; i < len(wallets); i++ {
		if wallets[i].TotalTxs != 0 {
			transmissionTypeRecords = buildTransmissionTypeRecords(buildRecords(wallets[i]))
			fmt.Print(transmissionTypeRecords)
		}

		transmissionType := models.TransmissionType{
			Name:    wallets[i].Name,
			Records: transmissionTypeRecords,
		}
		payload = append(payload, transmissionType)
	}

	if len(errors) > 0 {
		fmt.Print(errors)
	}

	//return payload, nil
	body, err := json.Marshal(payload)
	if err != nil {
		λ.MarshalError(err)
		return events.APIGatewayProxyResponse{}, err
	}
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: 200,
	}, nil
}


//func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
//	var payload []models.TransmissionType
//	//input
//	wallets, errors := DogeHandler()
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
//	if len(errors) > 0 {
//		fmt.Print(errors)
//	}
//
//	body, err := json.Marshal(payload)
//	if err != nil {
//		λ.MarshalError(err)
//	}
//	return events.APIGatewayProxyResponse{
//		Body:       string(body),
//		StatusCode: 200,
//	}, nil
//}

// loop over transmission_types
func buildRecords(wallet models.Wallet) []*models.Record {
	var records []*models.Record
	for i := 0; i < wallet.TotalTxs; i++ {
		record, err := cryptography.Decrypt(wallet.Txs[i].Incoming.Value, int64(i))
		if err.Status == 0 {
			records = append(records, record)
		}
	}
	return records
}
func buildTransmissionTypeRecords(records []*models.Record) models.TransmissionTypeRecords {
	fmt.Println(records)
	transmissionTypeRecords := records
	return transmissionTypeRecords
}

func main() {
	lambda.Start(Handler)
}

// swagger server code

//func GetRecordsHandlerFunc(params operations.GetRecordsParams) middleware.Responder {
//	return handle()
//}
//
//// loop over wallets and build records into each of them
//func handle() middleware.Responder {
//	var payload []models.TransmissionType
//	//input
//	wallets, errors := DogeHandler()
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
//	if len(errors) > 0 {
//		fmt.Print(errors)
//	}
//	return operations.NewGetRecordsOK().WithPayload(payload)
//}
