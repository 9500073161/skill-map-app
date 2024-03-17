package apis

import (
	"net/http"

	"github.com/9500073161/skill-map-app/models"
	"github.com/gin-gonic/gin"
)

func listUser(ctx *gin.Context) {

	user1 := models.User{
		"Prashob",
		"prashob.vp@outlook.com",
	}

	user2 := models.User{
		"itsforme",
		"itsforme@outlook.com",
	}
	ctx.JSON(http.StatusOK, []models.User{user1, user2})

}

func createUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

func updateUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}

func deleteUser(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})

}
