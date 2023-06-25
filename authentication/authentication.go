package authentication

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type AccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   string `json:"expires_in"`
}


func GetCredentials(consumerKey string, consumerSecret string) string {
	authString := consumerKey + ":" + consumerSecret
   base64AuthString := base64.StdEncoding.EncodeToString([]byte(authString))

   url := "https://api.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
   req, err := http.NewRequest("GET", url, nil)
   if err != nil {
	   fmt.Println("Error creating request:", err)
	   return ""
   }

   req.Header.Add("Authorization", "Basic "+base64AuthString)

   client := &http.Client{}
   res, err := client.Do(req)
   if err != nil {
	   fmt.Println("Error making request:", err)
	   return ""
   }
   defer res.Body.Close()

   body, err := io.ReadAll(res.Body)
   if err != nil {
	   fmt.Println("Error reading response body:", err)
	   return ""
   }

   // Parse JSON response
	var response AccessTokenResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON response:", err)
		return ""
	}

	accessToken := response.AccessToken
    return accessToken
}
