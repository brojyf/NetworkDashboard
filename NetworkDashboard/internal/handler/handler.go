package handler

import (
	"net/http"

	"github.com/brojyf/NetworkDashboard/internal/service"
	"github.com/gin-gonic/gin"
)

var queryService = service.NewQueryService()

func Handler(c *gin.Context) {
	category := c.Query("category")
	c.JSON(http.StatusOK, queryService.QueryByCategory(category))
}
