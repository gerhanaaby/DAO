package controllers

import (
	"DAO/models"
	"DAO/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostApplicant(c *gin.Context) {
	applicantRequest := models.Applicants{}

	if err := c.ShouldBindJSON(&applicantRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	//************** pindah foto  base64 -> simpan ke local storage -> buat/ambil url
	//  edit gin.Context KTPImage = URL, SelfieImage = URL

	err := services.CreateApplicant(applicantRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	//alamat

	//hit api history
	c.JSON(http.StatusCreated, gin.H{
		"message": "Yeay Berhasil!",
	})
}
