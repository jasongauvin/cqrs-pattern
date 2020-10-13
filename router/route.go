package router

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jasongauvin/cqrs_pattern/controllers"
)

func InitRouter() {
	r := gin.Default()
	r.POST("/users", controllers.CreateUser)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	fmt.Println("Listening to port :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
