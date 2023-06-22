package controllers

import (
	models "DAO/models/request"
	"DAO/services"
	"DAO/utils"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

// func PostApplicant1(c *gin.Context) {
// 	applicantRequest := models.SubmitApplicants{}
// 	KTPFile, errKTP := c.FormFile("KTPImage")
// 	if errKTP != nil {
// 		c.AbortWithError(http.StatusInternalServerError, errKTP)
// 		return
// 	}
// 	SelfieImageFile, errSelfie := c.FormFile("SelfieImage")
// 	if errSelfie != nil {
// 		c.AbortWithError(http.StatusInternalServerError, errSelfie)
// 		return
// 	}
// 	KTPPath := "images/KTP/" + KTPFile.Filename
// 	SelfiePath := "images/SelfieKTP/" + SelfieImageFile.Filename
// 	errorValidation := c.ShouldBind(&applicantRequest)
// 	if errorValidation != nil {
// 		c.String(http.StatusBadRequest, "error bad request")
// 		errorMessages := []string{}
// 		for _, e := range errorValidation.(validator.ValidationErrors) {
// 			errorMessage := fmt.Sprintf("error on field : %s, cditon: %s", e.Field(), e.ActualTag())
// 			errorMessages = append(errorMessages, errorMessage)
// 		}

// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": errorMessages,
// 		})
// 		return
// 	}
// 	errUri := c.ShouldBindUri(&applicantRequest)
// 	if errUri != nil {
// 		c.AbortWithError(http.StatusInternalServerError, errUri)
// 		return
// 	}
// 	errUploadKTP := c.SaveUploadedFile(applicantRequest.KTPImage, KTPPath)
// 	if errUploadKTP != nil {
// 		c.String(http.StatusInternalServerError, "error upload ktp image")
// 		return
// 	}
// 	errUploadSelfieKTP := c.SaveUploadedFile(applicantRequest.SelfieImage, SelfiePath)
// 	if errUploadSelfieKTP != nil {
// 		c.String(http.StatusInternalServerError, "error upload selfie image")
// 		return
// 	}
// 	errSendDataToDatabase := services.CreateApplicant(applicantRequest, KTPPath, SelfiePath)
// 	if errSendDataToDatabase != nil {
// 		c.AbortWithError(http.StatusInternalServerError, errSendDataToDatabase)
// 		return
// 	}

// 	c.JSON(http.StatusCreated, gin.H{
// 		"Message": "Successfully Submit Data Applicant!",
// 	})
// }

func PostApplicant(c *gin.Context) {
	// var validate *validator.Validate
	// validate = validator.New()
	applicantRequest := models.SubmitApplicants{}

	errorValidation := c.ShouldBindJSON(&applicantRequest)
	if errorValidation != nil {
		c.String(http.StatusBadRequest, "error bad request")
		errorMessages := []string{}
		for _, e := range errorValidation.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field : %s, cditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}
	KTPImage := utils.Base64toFile(applicantRequest.RawData.DataSubmit.KtpImage, "KTP", applicantRequest.CifCode)
	SelfieImage := utils.Base64toFile(applicantRequest.RawData.DataSubmit.SelfieImage, "SelfieImage", applicantRequest.CifCode)
	errSendDataToDatabase := services.CreateApplicant(applicantRequest, KTPImage, SelfieImage)
	if errSendDataToDatabase != nil {
		c.AbortWithError(http.StatusInternalServerError, errSendDataToDatabase)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Successfully Submit Data Applicant!",
	})
}

func GetApplicant(c *gin.Context) {

	services.GetApplicant()
	// if errGetDataFromDatabase != nil {
	// 	c.AbortWithError(http.StatusInternalServerError, errGetDataFromDatabase)
	// 	return
	// }

	// if getDataFromDatabase != nil {
	// 	c.JSON(http.StatusCreated, gin.H{
	// 		"Message": "Successfully Get Data Applicant!",
	// 		"Data":    getDataFromDatabase,
	// 	})
	// } else {
	// 	c.JSON(http.StatusCreated, gin.H{
	// 		"Message": "Successfully Get Data Applicant!",
	// 		"Data":    "No Data",
	// 	})
	// }
}
