package c2b

import (
	"fmt"
	"log"
	"main/authentication"
	"main/utils"
	"time"

	"github.com/go-resty/resty/v2"
)

func StkPush(phoneNumber string, amount float64, callbackURL string) string {
	timestamp := time.Now().Format("20060102150405")

	password := authentication.GenerateSTKPassword(
		utils.GoDotEnvVariable("shortCode"),
		utils.GoDotEnvVariable("passKey"),
		timestamp)
	fmt.Printf("callback url, %v", callbackURL)

	client := resty.New()

	consumerKey := utils.GoDotEnvVariable("consumerKey")
	consumerSecret := utils.GoDotEnvVariable("consumerSecret")
	accessToken := authentication.GetCredentials(consumerKey, consumerSecret)

	bearer := fmt.Sprintf("Bearer %v", accessToken)

	data := map[string]interface{}{
		"BusinessShortCode": utils.GoDotEnvVariable("shortCode"),
		"Password":          password,
		"Timestamp":         timestamp,
		"TransactionType":   "CustomerBuyGoodsOnline",
		"Amount":            amount,
		"PartyA":            phoneNumber,
		"PartyB":            "8186048",
		"PhoneNumber":       phoneNumber,
		"CallBackURL":       callbackURL,
		"AccountReference":  "AccountReference",
		"TransactionDesc":   "Transaction description",
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("Authorization", bearer).
		SetBody(data).
		Post("https://api.safaricom.co.ke/mpesa/stkpush/v1/processrequest")

	if err != nil {
		log.Fatal(err)
		return resp.String()
	}

	fmt.Println(resp.String())

	return resp.String()
}
