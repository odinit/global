package gweb

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	grpcapi "microbigdata/grpc_api/service"
	"net/http"
)

//c.Abort() 作为中间件时,不继续往下一个Handler执行,中断后返回
//c.Next() 继续向下一个Handler执行,后面的代码不继续执行

// AuthMiddle
// token验证相关,这里是作为中间件使用
// 在需要用户验证的请求handler前添加该中间件
func AuthMiddle(c *gin.Context) {
	// 获取token,token保存在Header中的Authorization字段中
	token := c.Request.Header.Get("Authorization")

	// 未获取token,直接返回异常
	if token == "" {
		zap.L().Error("token为空", zap.Any("request", c.Request.URL.String()))
		ErrTokenIsInvalid.Return(c)
		c.Abort()
		return
	}

	// 解析token
	auth, err := grpcapi.CheckLogin(token)
	if err != nil {
		zap.L().Error("grpcapi.CheckLogin(token)", zap.Error(err))
		ErrTokenParseFail.Return(c)
		c.Abort()
		return
	}

	// 如果用户为登录,返回异常
	if auth.CheckLogin == false {
		ErrUserNotLogin.Return(c)
		c.Abort()
		return
	}

	// 将该中间件获取的信息,传递到下一个Handler,避免重复操作
	c.Set("userInfo", auth)
	c.Next()
}

// Cors 处理跨域请求,支持options访问
func Cors(c *gin.Context) {
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
