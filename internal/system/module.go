package system

import "github.com/gofiber/fiber/v2"

type Module interface {
	Run()
}

type module struct {
	modules []Module
	router  *fiber.App
}

func WithModules(modules ...Module) func(*module) *module {
	return func(m *module) *module {
		m.modules = modules
		return m
	}
}

func WithRouter(router *fiber.App) func(*module) *module {
	return func(m *module) *module {
		m.router = router
		return m
	}
}

func NewModule(options ...func(*module) *module) module {
	var mod = module{}

	for _, option := range options {
		option(&mod)
	}

	return mod
}

func (m module) Start(appPort string) error {
	for _, mod := range m.modules {
		mod.Run()
	}

	return m.router.Listen(appPort)
}
