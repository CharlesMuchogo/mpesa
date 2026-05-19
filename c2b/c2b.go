package c2b

import (
	"fmt"
	"main/authentication"
	"main/utils"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/ochom/gutils/jsonx"
	"github.com/ochom/gutils/logs"
)

func StkPush(phoneNumber string, amount float64, callbackURL string) (map[string]any, error) {
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
		logs.Error("Calling mpesa callback response %+v", resp)
		return nil, err
	}

	return jsonx.Decode[map[string]any](resp.Body()), nil
}
