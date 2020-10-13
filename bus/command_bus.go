package bus

import (
	"fmt"

	"github.com/jasongauvin/cqrs_pattern/handlers"

	"github.com/jasongauvin/cqrs_pattern/structs"
)

func DispatchCommands(cmd *structs.Command) {
	switch cmd.Type {
	case "userCreationCommand":
		fmt.Println("Dispatch", cmd.Data)
		handlers.UserCreationCommand(cmd)
	}

}
