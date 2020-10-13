package main

import (
	"fmt"
	"log"

	"github.com/jasongauvin/cqrs_pattern/database"
	"github.com/jasongauvin/cqrs_pattern/router"

	"github.com/caarlos0/env/v6"
)

type config struct {
	// DbUser     string `env:"MYSQL_USER"`
	// DbPassword string `env:"MYSQL_PASSWORD"`
	// DbPort int `env:"DB_PORT" envDefault:"3306"`
	// DbHost     string `env:"DB_HOST"`
	// DbName     string `env:"DB_NAME"`
	DbHost     string `env:"DB_HOST"`
	DbName     string `env:"MYSQL_DATABASE"`
	DbUser     string `env:"MYSQL_USER"`
	DbPassword string `env:"MYSQL_PASSWORD"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	database.InitializeDb(cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbName)
	database.MakeMigrations()
	router.InitRouter()
	fmt.Println("suce")
}
