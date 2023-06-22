package controllers

import (
	models "DAO/models/request"
	"DAO/services"
	"DAO/utils"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCSR(c *gin.Context) {
	CSRRequest := models.CSR{}
	var errPos = ``

	requestBody := json.NewDecoder(c.Request.Body)
	err := requestBody.Decode(&CSRRequest)
	if err != nil {
		errPos = "Error Data Request"
	}

	JsonCSRByte, err := json.Marshal(CSRRequest)
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
		respone, err := services.ConsumeAPIService(utils.GetEndPoint("csr20", "csr20-key", "csr20-auth"), JsonCSRByte)
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
