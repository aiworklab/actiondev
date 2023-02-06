package api

import (
	"time"

	"github.com/aiworklab/actiondev/utils"
	"github.com/gin-gonic/gin"
)

// Api
type Api struct {
}

// NewApi
func NewApi() *Api {
	return &Api{}
}

// Ping godoc
// @Summary Ping
// @Description Ping
// @Tags Ping
// @Accept  json
// @Produce  json
// @Success 200 {string} string ""
// @Router /ping [get]
func (s *Api) Ping(c *gin.Context) {
	now := time.Now().Format("2006.01.02 03:04:05")
	utils.UpdateSystemStatus()
	c.JSON(200, gin.H{
		"message":    "pong",
		"date":       now,
		"sys_status": utils.SysStatus,
	})
}
