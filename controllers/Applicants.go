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

	applicantRequest := models.ApplicantsRequest{}



	KTPFile, errKTP := c.FormFile("KTPImage")
	SelfieImageFile, errSelfie := c.FormFile("SelfieImage")
	
	if errKTP != nil {
		c.AbortWithError(http.StatusInternalServerError, errKTP)
		return
	}
	
	if errSelfie != nil {
		c.AbortWithError(http.StatusInternalServerError, errSelfie)
		return
	}
	KTPPath := "images/KTP/" + KTPFile.Filename
	SelfiePath := "images/SelfieKTP/" + SelfieImageFile.Filename

	err := c.ShouldBind(&applicantRequest)
	if err != nil {
		//c.String(http.StatusInternalServerError, "error upload ktp image")
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field : %s, conditon: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}

		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}
	err1 := c.ShouldBindUri(&applicantRequest)
	if err1 != nil {
		c.AbortWithError(http.StatusInternalServerError, err1)
		return
	}
	err2 := c.SaveUploadedFile(applicantRequest.KTPImage, KTPPath)
	if err2 != nil {
		c.String(http.StatusInternalServerError, "error upload ktp image")
		return
	}
	err3 := c.SaveUploadedFile(applicantRequest.SelfieImage, SelfiePath)
	if err3 != nil {
		c.String(http.StatusInternalServerError, "error upload selfie image")
		return
	}

	err4 := services.CreateApplicant(applicantRequest, KTPPath, SelfiePath)
	if err4 != nil {
		c.AbortWithError(http.StatusInternalServerError, err2)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Successfully Submit Data Applicant!",
	})
}
