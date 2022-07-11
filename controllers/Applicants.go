package controllers

import (
	"DAO/models"
	"DAO/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func PostApplicant(c *gin.Context) {

	applicantRequest := models.Applicants{}
	//applicantRespone := models.Applicants{}

	//applicantHistoryRequest := models.ApplicantHistory{}

	err1 := c.ShouldBindJSON(&applicantRequest)
	if err1 != nil {
		errorMessages := []string{}
		for _, e := range err1.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field : %s, conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	applicantRespone, err2 := services.CreateApplicant(applicantRequest)
	if err2 != nil {
		c.AbortWithError(http.StatusInternalServerError, err2)
		return
	}

	// err3 := services.CreateApplicantHistory(applicantHistoryRequest)
	// if err3 != nil {
	// 	c.AbortWithError(http.StatusInternalServerError, err3)
	// 	return
	// }

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully Submit Data Applicant!",
		"data":    applicantRespone,
		//"history applicant": applicantHistoryRequest,
	})
}
