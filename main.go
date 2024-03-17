package main

import (
	"fmt"

	"github.com/9500073161/skill-map-app/apis"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Developing Skill-MAP")

	router := gin.Default()

	apis.RegisterUserApis(router)

	router.Run()

}
