package controllers

import (
	"DAO/models"
	"DAO/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCustomer(c *gin.Context) {
	customerRequest := models.Customer{}

	if err := c.ShouldBindJSON(&customerRequest); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := services.CreateCustomer(customerRequest)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Yeay Berhasil!",
	})
}
