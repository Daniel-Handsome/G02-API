package routes

import (
	"G02-API/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func RegisterApiRoutes(r *gin.Engine) {
	// group can not use prefix name
	v1 := r.Group("v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "ok",
			})
		})
		authGroup := v1.Group("auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			// 判断 Email 是否已注册
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			// 发送验证码
			vcc := new(auth.VerifyCodeController)
			// 图片验证码，需要加限流
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		}
	}
}

func Register404Handler(r *gin.Engine) {
	r.NoRoute(func(context *gin.Context) {
		acceptStrin := context.Request.Header.Get("Accept")
		if strings.Contains(acceptStrin, "text/html") {
			context.String(http.StatusNotFound, "not found page")
		} else {
			context.JSON(http.StatusNotFound, "page not found")
		}
	})
}
