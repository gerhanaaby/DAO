package controllers

import (
	"DAO/models"
	"DAO/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostApplicantHistory(c *gin.Context) {
	applicantRequest := models.ApplicantHistory{}

	if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := services.CreateApplicantHistory(applicantRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//alamat
	c.JSON(http.StatusCreated, gin.H{
		"message": "Yeay Berhasil! History",
	})
}

func PutApplicantHistory(c *gin.Context) {
	applicantRequest := models.ApplicantHistory{}

	if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := services.UpdateApplicantHistory(c.Param("ApplicantID"), applicantRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//alamat
	c.JSON(http.StatusCreated, gin.H{
		"message": "Yeay Berhasil! History",
	})
}
