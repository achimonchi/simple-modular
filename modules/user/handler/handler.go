package handler

import (
	"context"
	"simple-modular/modules/user/domain"
)

// handler membutuhkan service
// jadi contract di declare di handler
type service interface {
	GetAll() (users []domain.User, err error)
	Create(ctx context.Context, req domain.User) (err error)
}

type handler struct {
	svc service
}

func NewHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}
