package user

import (
	"log"
	"simple-modular/modules/user/domainhandler"
	"simple-modular/modules/user/internal/handler"
	"simple-modular/modules/user/internal/repository"
	"simple-modular/modules/user/internal/service"

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

	u.RegisterInternalHandler()

	repo := repository.NewRepository()
	svc := service.NewService(repo)
	handler := handler.NewHandler(svc)

	userRoute := u.router.Group("users")
	{
		userRoute.Get("/", handler.GetAll)
		userRoute.Post("/", handler.CreateNewUser)
	}
}

func (u userModule) RegisterInternalHandler() {
	repo := repository.NewRepository()
	svc := service.NewService(repo)

	domainhandler.InitHandler(svc)

}
