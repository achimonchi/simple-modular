package handler

import (
	"context"
	"simple-modular/modules/transaction/domain"
)

type service interface {
	Transfer(ctx context.Context, req domain.Transaction) (err error)
	Topup(ctx context.Context, req domain.Transaction) (err error)
}

type handler struct {
	svc service
}

func NewHandler(svc service) handler {
	return handler{
		svc: svc,
	}
}
