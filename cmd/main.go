package main

import (
	"fmt"

	"github.com/Anvar571/golang-crud/configs"
	"github.com/Anvar571/golang-crud/storage/postgres"
)

func main() {
	cfg := configs.Load()

	fmt.Println(cfg)

	sqlDB, err := postgres.ConnectionDb(&cfg)
	if err != nil {
		panic(err)
	}

	fmt.Println(sqlDB)

	fmt.Println("importing config file")
}
