package apis

import (
	"github.com/gin-gonic/gin"
)

func RegisterUserApis(router *gin.Engine) {

	userGroup := router.Group("api/users/")

	userGroup.GET("", listUser)

	userGroup.POST("", createUser)

	userGroup.PATCH("", updateUser)

	userGroup.DELETE("", deleteUser)

}
