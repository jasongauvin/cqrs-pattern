package services

import (
	"fmt"
	"net/http"

	"github.com/jasongauvin/cqrs_pattern/bus"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/cqrs_pattern/structs"
)

func CreateUserCommand(c *gin.Context) {
	fmt.Println("Request", c.Request)
	var userForm structs.UserForm

	// Si json ne correspond pas a ma donnée
	if err := c.ShouldBindJSON(&userForm); err != nil {
		c.JSON(http.StatusBadRequest, err.Error)
		c.Abort()
		return
	}
	fmt.Println("Services", userForm)
	// TODO : Si header ne correspond au bon content-type

	var command structs.Command

	//hydrater la donnée
	//TODO = Dossier des constantes globale en appelant le constante container
	command.Type = "userCreationCommand"
	command.Data = userForm

	bus.DispatchCommands(&command)

}
