package controllers

import (
	models "DAO/models/request"
	"DAO/services"
	"DAO/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostSyncCIF(c *gin.Context) {
	SyncCIFRequest := models.Sync{}
	var errPos = ``

	requestBody := json.NewDecoder(c.Request.Body)
	err := requestBody.Decode(&SyncCIFRequest)
	if err != nil {
		errPos = "Error Data Request"
	}

	JsonSyncCIFByte, err := json.Marshal(SyncCIFRequest)
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
		respone, err := services.ConsumeAPIService(utils.GetEndPoint("synccif", "synccif-key", "synccif-auth"), JsonSyncCIFByte)
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
