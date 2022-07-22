package controllers

import (
	"DAO/models"
	"DAO/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
)

func PostApplicant(c *gin.Context) {

	applicantRequest := models.ApplicantsRequest{}

	KTPFile, errKTP := c.FormFile("KTPImage")

	if errKTP != nil {
		c.AbortWithError(http.StatusInternalServerError, errKTP)
		return
	}
	SelfieImageFile, errSelfie := c.FormFile("SelfieImage")
	if errSelfie != nil {
		c.AbortWithError(http.StatusInternalServerError, errSelfie)
		return
	}
	
	KTPPath := "images/KTP/" + KTPFile.Filename
	SelfiePath := "images/SelfieKTP/" + SelfieImageFile.Filename
	
	err := c.c.PostFormMap(&applicantRequest)
	if err != nil {
		//c.String(http.StatusInteralServerError, "error upload ktp image")
		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("error on field : %s, cnditon: %s", e.Field(), e.ActualTag())
			rrorMessages = append(errorMessages, errorMessage)
	}

		c.JSON(http.StatusBadReqest, gin.H{
			"rror": errorMessages,
		})
		eturn
	}
	err1 := c.ShouldindUri(&applicantRequest)
	if err1 != nil {
		c.AborWithError(http.StatusInternalServerError, err1)
		eturn
	}
	err2 := c.SaveUpoadedFile(applicantRequest.KTPImage, KTPPath)
	if err2 != nil {
		c.Strig(http.StatusInternalServerError, "error upload ktp image")
		eturn
	}
	err3 := c.SaveUpoadedFile(applicantRequest.SelfieImage, SelfiePath)
	if err3 != nil {
		c.Strig(http.StatusInternalServerError, "error upload selfie image")
		eturn
}

	err4 := servicesCreateApplicant(applicantRequest, KTPPath, SelfiePath)
	if err4 != nil {
		c.AborWithError(http.StatusInternalServerError, err2)
		eturn
}

	c.JSON(http.StatusCreated, gin.H{
		"Message": "Successfully Submit Data Applicant!",
})
}
