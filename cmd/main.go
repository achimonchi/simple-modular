package main

import (
	"simple-modular/internal/system"
	"simple-modular/modules/transaction"
	"simple-modular/modules/user"

	"github.com/gofiber/fiber/v2"
)

func main() {

	router := fiber.New(fiber.Config{
		Prefork: true,
	})

	var userModule = user.NewUserModule(router)
	var transactionModule = transaction.NewtransactionModule(router)

	var module = system.NewModule(
		system.WithRouter(router),
		system.WithModules(
			userModule,
			transactionModule,
		),
	)

	if err := module.Start(":4444"); err != nil {
		panic(err)
	}
}
