package routes

import (
	"DAO/controllers"

	"github.com/gin-gonic/gin"
)

func CallRoutes(port string) {
	router := gin.Default()
	v1 := router.Group("/v1/")
	v1.POST("/submitapplicant", controllers.PostApplicant)
	v1.POST("/submithistory", controllers.PostApplicantHistory)
	v1.POST("/screeningOneFCC", controllers.PostScreening)
	v1.POST("/syncCIF", controllers.PostSyncCIF)
	v1.POST("/CSR", controllers.PostCSR)
	v1.PUT("/updatehistory/{ApplicantID}", controllers.PutApplicantHistory)
	v1.GET("/getData", controllers.GetApplicant)
	router.Run(port)
}
