package Handler

import (
	"github.com/brojyf/NetworkDashboard/internal/model"
	"github.com/gin-gonic/gin"
)

var dataMap = map[string]interface{}{
	"ai":              model.MockData.AI,
	"search_engine":   model.MockData.SearchEngine,
	"cdn":             model.MockData.CDN,
	"social":          model.MockData.Social,
	"cloud":           model.MockData.Cloud,
	"video_streaming": model.MockData.VideoStreaming,
}

func Handler(c *gin.Context) {
	category := c.Query("category")

	if data, found := dataMap[category]; found {
		c.JSON(200, data)
		return
	}
	c.JSON(400, gin.H{"error": "category not found"})
}
