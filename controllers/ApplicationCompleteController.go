package controllers

import (
	models "DAO/models/request"
	"DAO/services"
	"DAO/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostApplicationComplete(c *gin.Context) {
	AppCompleteRequest := models.ApplicationComplete{}
	var errPos = ``

	requestBody := json.NewDecoder(c.Request.Body)
	err := requestBody.Decode(&AppCompleteRequest)
	if err != nil {
		errPos = "Error Data Request"
	}

	JsonApplicationCompleteByte, err := json.Marshal(AppCompleteRequest)
	if err != nil {
		errPos = "Error Data Request"
		return
	}

	if errPos != `` {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errPos,
		})
		return
	} else {
		respone, err := services.ConsumeAPIService(utils.GetEndPoint("eformcentral", "eformcentral-key", "eformcentral-auth"), JsonApplicationCompleteByte)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		var data map[string]interface{}
		if err := json.Unmarshal(respone, &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": err,
			})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "success screening data",
			"data":    data,
		})

	}
}
