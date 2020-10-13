package repository

import (
	"fmt"

	"github.com/jasongauvin/cqrs_pattern/models"
)

func CreateUserToDatabase(user models.User) {
	//logique d'insertion en db
	fmt.Println("Repository", user)

	// result := database.DB.Create(user) // pass pointer of data to Create

	// fmt.Println("user", user.ID)                     // returns inserted data's primary key
	// fmt.Println("error", result.Error) // returns error
	// fmt.Println("row affected", result.RowsAffected) // returns inserted records count
}
