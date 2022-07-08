package routes

import (
	"DAO/controllers"

	"github.com/gin-gonic/gin"
)

func CallRoutes() {
	router := gin.Default()
	v1 := router.Group("/v1")
	v1.POST("/submitapplicant", controllers.PostApplicant)
	v1.POST("/submithistory", controllers.PostApplicantHistory)
	v1.PUT("/updatehistory/{ApplicantID}", controllers.PutApplicantHistory)
}
