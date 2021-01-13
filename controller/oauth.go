// 只实现OAuth2中的Password模式作系统登陆用途，其他没有用到的细节均忽略
package controller

import (
	"github.com/fenixsoft/monolithic_arch_golang/infrasturcture"
	"github.com/fenixsoft/monolithic_arch_golang/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type OAuth2Controller struct{}

func (c *OAuth2Controller) Register(router gin.IRouter) {
	router.GET("/oauth/token", func(context *gin.Context) {
		// 验证用户身份
		username := context.Query("username")
		password := context.Query("password")
		s := service.Service{Context: context.Request.Context()}
		if account, err := s.CheckUserAccount(username, password); err != nil {
			infrasturcture.ClientError(context, "用户名或密码不正确")
		} else {
			context.JSON(http.StatusOK, s.BuildJWTAccessToken(account))
		}
	})
}