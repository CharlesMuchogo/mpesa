package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CallbackUrl(context *gin.Context) {

	body, err := context.GetRawData()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read request body"})
		return
	}

	fmt.Println(string(body))
	context.JSON(http.StatusOK, gin.H{"message": "Received data"})
}
