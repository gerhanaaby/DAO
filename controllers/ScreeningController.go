package controllers

import (
	"DAO/models/request"
	"DAO/services"
	"DAO/utils"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostScreening(c *gin.Context) {
	screeningRequest := models.Screening{}

	var errPos = ``

	requestBody := json.NewDecoder(c.Request.Body)
	err := requestBody.Decode(&screeningRequest)
	if err != nil {
		errPos = "Error Data Request"
	}

	JsonOneFCCByte, err := json.Marshal(screeningRequest)
	if err != nil {
		errPos = "Error Data Request"
		return
	}

	MatchScores, _ := strconv.Atoi(screeningRequest.MatchScore)
	if MatchScores <= 75 {
		errPos = "App Rejected"
		return
	}

	if errPos != `` {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": errPos,
		})
		return
	} else {
		respone, err := services.ConsumeAPIService(utils.GetEndPoint("onefcc", "onefcc-key", "onefcc-auth"), JsonOneFCCByte)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "success screening data",
			"data":    json.Unmarshal(respone, &screeningRequest),
		})

	}

}
