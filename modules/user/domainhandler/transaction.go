package domainhandler

import (
	"context"
	"simple-modular/modules/user/entity"
)

// FindUserById implements adapter.userRepository.
func (h *handler) FindUserById(ctx context.Context, userId int) (user entity.User, err error) {
	userDomain, err := h.svc.FindUserById(ctx, userId)
	if err != nil {
		return
	}

	user = entity.User{
		Id:      userDomain.Id,
		Name:    userDomain.Name,
		Balance: userDomain.Balance,
	}
	return
}

// UpdateBalance implements adapter.userRepository.
func (h *handler) UpdateBalance(ctx context.Context, userId int, balance int) (err error) {
	return h.svc.UpdateBalance(ctx, userId, balance)
}
