package c2b

import (
	"fmt"
	"log"
	"main/authentication"
	"main/utils"
	"time"

	"github.com/go-resty/resty/v2"
)

func StkPush(phoneNumber string, amount float64, callbackURL string) {
	timestamp := time.Now().Format("20060102150405")

	password := authentication.GenerateSTKPassword(
		utils.GoDotEnvVariable("shortCode"),
		utils.GoDotEnvVariable("passKey"),
		timestamp)


	client := resty.New()

	consumerKey := utils.GoDotEnvVariable("consumerKey")
	consumerSecret := utils.GoDotEnvVariable("consumerSecret")
    accessToken :=  authentication.GetCredentials(consumerKey , consumerSecret )

	bearer := fmt.Sprintf("Bearer %v", accessToken)

	data := map[string]interface{}{ 
		"BusinessShortCode": utils.GoDotEnvVariable("shortCode"),
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerBuyGoodsOnline",
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            utils.GoDotEnvVariable("shortCode"),
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       callbackURL,
		"AccountReference":  "mpesa",
		"TransactionDesc":   "Transaction description",
	}


	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", bearer).
		SetBody(data).
		Post("https://api.safaricom.co.ke/mpesa/stkpush/v1/processrequest")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp.String())
}