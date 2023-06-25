package main

import (
	"fmt"
	"main/authentication"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	
	consumerKey := goDotEnvVariable("consumerKey")
	consumerSecret := goDotEnvVariable("consumerSecret")
   accessToken :=  authentication.GetCredentials(consumerKey , consumerSecret )
   fmt.Println("Access Token:", accessToken)
     
 
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	CheckError(err)

	return os.Getenv(key)
}

func CheckError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

