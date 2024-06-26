package domainhandler

import (
	"context"
	"simple-modular/modules/user/internal/domain"
)

type handler struct {
	svc service
}

type service interface {
	FindUserById(ctx context.Context, userId int) (user domain.User, err error)
	UpdateBalance(ctx context.Context, userId int, balance int) (err error)
}

var h *handler

func InitHandler(svc service) {
	h = &handler{
		svc: svc,
	}
}

func GetHandler() *handler {
	return h
}
