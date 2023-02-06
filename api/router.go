package api

import (
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine) {
	api := NewApi()
	r.GET("/ping", api.Ping)
}
