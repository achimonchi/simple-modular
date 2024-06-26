package service

import (
	"context"
	"simple-modular/modules/user/internal/domain"
)

func (s service) GetAll() (users []domain.User, err error) {
	return s.repo.FindAll()
}

// FindUserById implements domainhandler.service.
func (s service) FindUserById(ctx context.Context, userId int) (user domain.User, err error) {
	return s.repo.FindUserById(ctx, userId)
}
