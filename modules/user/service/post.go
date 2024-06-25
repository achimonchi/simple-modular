package service

import (
	"context"
	"simple-modular/modules/user/domain"
)

func (s service) Create(ctx context.Context, request domain.User) (err error) {
	return s.repo.Insert(ctx, request)
}
