package routes

import (
	"leiloa/controllers"
	"leiloa/middleware"

	"github.com/gin-gonic/gin"
)

func StartRoutes(r *gin.Engine) {
	r.POST("/api/user/create", controllers.CreateUser)
	r.POST("/api/proposal/create", middleware.RequireAuth, controllers.CreateProposal)
	r.POST("/api/auction/create", middleware.RequireAuth, controllers.CreateAuction)

}
