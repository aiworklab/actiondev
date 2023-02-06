package utils

import (
	"net/http"

	"github.com/aichy126/igo"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Rfail 错误返回
func Rfail(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": 1,
		"msg":  msg,
		"data": nil,
	})
}

// Rsucc 成功返回
func Rsucc(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "succeed",
	})
}

func ConfGetbool(path string) bool {
	return igo.App.Conf.GetBool(path)
}

func ConfGetString(path string) string {
	return igo.App.Conf.GetString(path)
}

func ConfGetInt(path string) int {
	return igo.App.Conf.GetInt(path)
}

func ConfGetInt64(path string) int64 {
	return igo.App.Conf.GetInt64(path)
}

func GenerateToken() string {
	uid, _ := uuid.NewUUID()
	return uid.String()
}
