package Controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoadPriPage 加载最初界面接口
// @Summary 加载最初界面接口
// @Description 加载最初界面
// @Tags 加载最初界面接口
// @Security ApiKeyAuth
// @Success 200
// @Router /pri/ [get]
func LoadPriPage() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "page load success",
		})
	}
}
