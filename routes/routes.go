package routes

import (
	pointshandlers "bitresume/api/pointsHandlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.POST("/api/points_logs/",pointshandlers.HandlePointlogs)
	r.POST("/api/activity_graph/",)
}