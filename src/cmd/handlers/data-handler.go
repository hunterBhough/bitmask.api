package handlers

// Fill the wallet struct with wallet names and addresses
// Dependency: main function
// dumb component

import (
	//"fmt"
	//"os"
	//"encoding/json"
	//"io/ioutil"
	"github.com/hunterBhough/go-doge/src/models"
)

type ResponseType struct {
	Wallets []models.Wallet
}

func TransmuteJSON() ResponseType {
	var response ResponseType
	var wallets []models.Wallet
	foodAddress := models.Data{Address: "DSnFEMUgqi1Y6x3akQFYX8az8KPa8deAnU"}
	airAddress := models.Data{Address: "DKZddE5XP6uZTQXUegaqpPhxZEbeYkfPkK"}
	wallets = append(wallets, models.Wallet{
		Name: "NNDSS Food borne",
		Data: foodAddress,
	})
	wallets = append(wallets, models.Wallet{
		Name: "NNDSS Airborne",
		Data: airAddress,
	})
	//file, e := ioutil.ReadFile("/Users/hhough/code/go/src/github.com/hunterBhough/go-doge/data/database.json")
	//if e != nil {
	//	fmt.Printf("File error: %v\n", e)
	//	os.Exit(1)
	//}
	//
	////m := new(Dispatch)
	////var m interface{}
	//e = json.Unmarshal(file, &response)
	//if e != nil {
	//	fmt.Printf("Unmarshal error: %v\n", e)
	//	os.Exit(1)
	//}
	response = ResponseType{wallets}
	return response
}
