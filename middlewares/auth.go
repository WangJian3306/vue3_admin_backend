package middlewares

import (
	"vue3_admin/controller"
	"vue3_admin/pkg/jwt"

	"github.com/gin-gonic/gin"
)

// JWTAuthMiddleware 基于JWT的认证中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {
		// 客户端携带 Token 有三种方式 1.放在请求头 2.放在请求体 3.放在URI
		// 这里假设Token放在Header的Authorization中，并使用Bearer开头
		// Authorization: Bearer xxxxx.xxx.xxx
		// 这里的具体实现方式要依据你的实际业务情况决定
		//authHeader := c.Request.Header.Get("Authorization")
		//if authHeader == "" {
		//	controller.ResponseError(c, controller.CodeNeedLogin)
		//	c.Abort()
		//	return
		//}
		//// 按空格分割
		//parts := strings.SplitN(authHeader, " ", 2)
		//if !(len(parts) == 2 && parts[0] == "Bearer") {
		//	controller.ResponseError(c, controller.CodeInvalidToken)
		//	c.Abort()
		//	return
		//}
		//// parts[1]是获取到的tokenString，我们使用之前定义好的解析JWT的函数来解析它
		//mc, err := jwt.ParseToken(parts[1])
		//if err != nil {
		//	controller.ResponseError(c, controller.CodeInvalidToken)
		//	c.Abort()
		//	return
		//}

		// 获取Header中的Authorization
		token := c.Request.Header.Get("Token")
		if token == "" {
			controller.ResponseError(c, controller.CodeNeedLogin)
			c.Abort()
			return
		}

		// 解析 Token
		mc, err := jwt.ParseToken(token)
		if err != nil {
			controller.ResponseError(c, controller.CodeInvalidToken)
			c.Abort()
			return
		}

		// 将当前请求的userID信息保存到请求的上下文 c 上
		c.Set(controller.CtxUserIDKey, mc.UserID)
		c.Next() // 后续的处理请求的函数中，可以使用c.Get(CtxUserIDKey)来获取当前请求的用户信息
	}
}
