
package routes

import (
	"backend-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		applicants := api.Group("/applicants")
		{
			applicants.POST("", controllers.CreateApplicant)
			applicants.GET("", controllers.GetApplicants)
			applicants.PATCH("/:id/stage", controllers.UpdateApplicantStage)
		}
	}
}
