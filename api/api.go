package api

import (
	"fmt"
	"main/c2b"
	"main/structs"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ochom/gutils/helpers"
)

func MpesaExpress(c *gin.Context) {
	var request structs.MpesaExpress

	if err := c.ShouldBindJSON(&request); err != nil {
		fmt.Printf("error: %s \n ", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	phone, ok := helpers.ParseMobile(request.Phone)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid phone number"})
		return
	}

	response, err := c2b.StkPush(phone, request.Amount, "https://charlesmuchogo.com/api/callback")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, response)
}
