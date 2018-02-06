package handlers

// Get a list of transactions for each wallet
// Dependency: Wallet generation in data-handler.go
// smart component

import (
	"github.com/hunterBhough/go-doge/src/models"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

var network = "DOGE"

func DogeHandler() ([]models.Wallet, []models.Error) {
	dat := TransmuteJSON()
	wallets, errors := getWalletDetails(dat.Wallets)
	return wallets, errors
}

func getWalletDetails(wallets []models.Wallet) ([]models.Wallet, []models.Error) {
	var payload []models.Wallet
	var errors []models.Error

	for i := 0; i < len(wallets); i++ {
		builtURL := buildAddressUrl(network, wallets[i].Address)
		response, err := fillWallet(builtURL)
		if err.Status == 0 {
			wallet := models.Wallet{
				Name: wallets[i].Name,
				Data: response.Data,
			}
			payload = append(payload, wallet)
		} else {
			errors = append(errors, err)
		}
	}
	return payload, errors
}

func buildAddressUrl(network string, address string) string {
	u := "https://chain.so/api/v2/address/" + network + "/" + address
	return u
}

func fillWallet(s string) (models.Wallet, models.Error) {
	var err models.Error
	var response models.Wallet

	resp, e := http.Get(s)
	if e != nil {
		err = models.Error{
			Status:  1,
			Message: "failed to get response from Doge",
		}
	}
	defer resp.Body.Close()
	body, e := ioutil.ReadAll(resp.Body)
	e = json.Unmarshal(body, &response)
	if e != nil {
		err = models.Error{
			Status:  2,
			Message: "failed to get unmarshal response",
		}
	}
	return response, err
}
