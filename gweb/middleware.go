package gweb

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//c.Abort() 作为中间件时,不继续往下一个Handler执行,中断后返回
//c.Next() 继续向下一个Handler执行,后面的代码不继续执行

// CSRFMiddle 处理跨域请求,支持options访问
func CSRFMiddle(c *gin.Context) {
	method := c.Request.Method

	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
	c.Header("Access-Control-Allow-Credentials", "true")

	//放行所有OPTIONS方法
	if method == "OPTIONS" {
		c.AbortWithStatus(http.StatusNoContent)
	}
	// 处理请求
	c.Next()
}
