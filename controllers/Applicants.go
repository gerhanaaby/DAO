package controllers

import (
	"DAO/models"
	"DAO/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostApplicant(c *gin.Context) {
	applicantRequest := models.ApplicantsRequest{}

	if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := services.CreateApplicant(applicantRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Yeay Berhasil!",
	})
}
