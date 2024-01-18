package api

import (
	"fmt"
	"main/c2b"
	"main/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func MpesaExpress(c *gin.Context) {
	var mpesa_express structs.MpesaExpress

	if err := c.ShouldBindJSON(&mpesa_express); err != nil {
		fmt.Printf("error: %s \n ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	response := c2b.StkPush(mpesa_express.Phone, mpesa_express.Amount, "https://patakazi.com/pay/callback")
	c.JSON(http.StatusOK, gin.H{"message": response})
}
