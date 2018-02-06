package handlers

// Fill the wallet struct with wallet names and addresses
// Dependency: main function
// dumb component

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"github.com/hunterBhough/go-doge/src/models"
)

type ResponseType struct {
	Wallets []models.Wallet
}

// Main function
// I realize this function is much too simple I am simply at a loss to

func TransmuteJSON() ResponseType {
	var response ResponseType

	file, e := ioutil.ReadFile("/Users/hhough/code/go/src/github.com/hunterBhough/go-doge/data/database.json")
	if e != nil {
		fmt.Printf("File error: %v\n", e)
		os.Exit(1)
	}

	//m := new(Dispatch)
	//var m interface{}
	e = json.Unmarshal(file, &response)
	if e != nil {
		fmt.Printf("Unmarshal error: %v\n", e)
		os.Exit(1)
	}
	return response
}
