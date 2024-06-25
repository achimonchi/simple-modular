package user

import (
	"log"
	"simple-modular/modules/user/handler"
	"simple-modular/modules/user/repository"
	"simple-modular/modules/user/service"

	"github.com/gofiber/fiber/v2"
)

type userModule struct {
	router fiber.Router
}

func NewUserModule(router fiber.Router) userModule {
	return userModule{
		router: router,
	}
}

func (u userModule) Run() {
	log.Println("Running user module")

	repo := repository.NewRepository()
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	userRoute := u.router.Group("users")
	{
		userRoute.Get("/", handler.GetAll)
		userRoute.Post("/", handler.CreateNewUser)
	}
}
