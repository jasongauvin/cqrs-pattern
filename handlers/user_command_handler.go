package handlers

import (
	"fmt"

	"github.com/jasongauvin/cqrs_pattern/repository"

	"github.com/jasongauvin/cqrs_pattern/models"
	"github.com/jasongauvin/cqrs_pattern/structs"
)

func UserCreationCommand(cmd *structs.Command) {
	var user models.User
	userForm := cmd.Data.(structs.UserForm)
	//Je reafecte les valeurs pour binder
	user.Email = userForm.Email
	user.Name = userForm.Name
	fmt.Println("Handler", user)
	repository.CreateUserToDatabase(user)
}
