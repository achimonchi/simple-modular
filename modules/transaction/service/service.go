package service

import (
	"context"
	"simple-modular/modules/transaction/domain"
)

type service struct {
	repo repository
	user userAdapter
}

func NewService(repo repository) service {
	return service{
		repo: repo,
	}
}

func (s *service) WithUserAdapter(adapter userAdapter) {
	s.user = adapter
}

// service butuh repository sebagai perantara untuk data flow
// jadi kontrak di define di service/consumer
type repository interface {
	Insert(ctx context.Context, data domain.Transaction) (err error)
	Topup(ctx context.Context, data domain.Transaction) (err error)
}

type userAdapter interface {
	FindUserById(ctx context.Context, userId int) (user domain.User, err error)
	UpdateBalance(ctx context.Context, userId int, balance int) (err error)
}
