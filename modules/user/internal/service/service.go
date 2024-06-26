package service

import (
	"context"
	"simple-modular/modules/user/internal/domain"
)

type service struct {
	repo repository
}

func NewService(repo repository) service {
	return service{
		repo: repo,
	}
}

// service butuh repository sebagai perantara untuk data flow
// jadi kontrak di define di service/consumer
type repository interface {
	FindAll() (users []domain.User, err error)
	FindUserById(ctx context.Context, userId int) (user domain.User, err error)
	Insert(ctx context.Context, user domain.User) (err error)
	UpdateBalance(ctx context.Context, userId int, balance int) (err error)
}
