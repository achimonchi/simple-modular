package transaction

import (
	"log"
	"simple-modular/modules/transaction/adapter"
	"simple-modular/modules/transaction/handler"
	"simple-modular/modules/transaction/repository"
	"simple-modular/modules/transaction/service"
	userRepository "simple-modular/modules/user/repository"

	"github.com/gofiber/fiber/v2"
)

type transactionModule struct {
	router fiber.Router
}

func NewtransactionModule(router fiber.Router) transactionModule {
	return transactionModule{
		router: router,
	}
}

func (t transactionModule) Run() {
	log.Println("Running transaction module")
	repo := repository.NewRepository()
	userRepo := userRepository.NewRepository()
	userAdapter := adapter.NewUserAdapter(userRepo)

	svc := service.NewService(repo)
	svc.WithUserAdapter(userAdapter)
	handler := handler.NewHandler(svc)

	trxRoute := t.router.Group("transactions")
	{
		trxRoute.Post("/transfer", handler.Transfer)
		trxRoute.Post("/topup", handler.Topup)
	}
}
