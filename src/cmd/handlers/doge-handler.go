package handlers

import (
	"github.com/hunterBhough/go-doge/src/models"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"net/http"
)

var network = "DOGE"

// for each wallet
// for each transaction
//build transactions into

func DogeHandler() []models.Wallet {
	dat := TransmuteJSON()
	//wallets := convertWallets(storeTransactions(dat.Wallets))
	wallets := getWalletDetails(dat.Wallets)
	return wallets
}

func getWalletDetails(wallets []models.Wallet) []models.Wallet {
	payload := []models.Wallet{}

	for i := 0; i < len(wallets); i++ {
		builtURL := buildAddressUrl(network, wallets[i].Address)
		response := fillWallet(builtURL)
		wallet := models.Wallet{
			Name: wallets[i].Name,
			Data: response.Data,
		}
		//fmt.Println(response.Data.Txs[0].Incoming.Value)
		payload = append(payload, wallet)
	}
	return payload
}

func buildAddressUrl(network string, address string) string {
	u := "https://chain.so/api/v2/address/" + network + "/" + address
	return u
}

func fillWallet(s string) models.Wallet {
	resp, err := http.Get(s)
	if err != nil {
		// handle error
		fmt.Println("error:", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	var response models.Wallet
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("error:", err)
	}
	return response
}




//func convertWallets(wallets []models.DogeWalletResponse) []models.Wallet {
//	payload := []models.Wallet{}
//
//	for i := 0; i < len(wallets); i++ {
//		wallet := models.Wallet{
//			Transactions: convertTransactions(wallets[i]),
//		}
//		payload = append(payload, wallet)
//	}
//	return payload
//}
//
//func convertTransactions(wallet models.DogeWalletResponse) models.WalletTransactions {
//	payload := []models.Transaction{}
//	newPayload := models.WalletTransactions{}
//	for j := 0; j < len(wallet.Data.Txs); j++ {
//		value, err := strconv.ParseFloat(wallet.Data.Txs[j].Incoming.Value, 64)
//		if err != nil {
//			// todo handle error
//			fmt.Println("error:", err)
//		} else {
//			transaction := &models.Transaction{
//				Amount: value,
//				Hash: wallet.Data.Txs[j].Txid,
//				Time: int64(wallet.Data.Txs[j].Time),
//			}
//			payload = append(payload, *transaction)
//			newPayload = models.WalletTransactions{
//				transaction,
//			}
//		}
//	}
//	newPayload = models.WalletTransactions{
//		*transaction
//	}
//	return newPayload
//
//}