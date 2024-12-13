package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理跨域请求，支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求方法
		method := c.Request.Method

		// 添加跨域响应头
		c.Header("Content-type", "application/json")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Max-Age", "86400")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Token, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Header("Access-Control-Allow-Credentials", "false")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		// 处理请求
		c.Next()
	}
}
