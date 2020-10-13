package controllers

import (
	"fmt"

	"github.com/jasongauvin/cqrs_pattern/services"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	fmt.Println("UserController")
	services.CreateUserCommand(c)
}
