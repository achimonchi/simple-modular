package adapter

import (
	"context"
	"simple-modular/modules/transaction/domain"
	userDomain "simple-modular/modules/user/domain"
)

// depend to userDomain
type userRepository interface {
	FindUserById(ctx context.Context, userId int) (user userDomain.User, err error)
	UpdateBalance(ctx context.Context, userId int, balance int) (err error)
}

// untuk interaksi antar cross domain, kita butuh "adapter"
// adapter ini akan bertugas sebagai proxy ke domain lain
// tujuannya adalah agar saat terjadinya cara komunikasi antar domain, maka cukup mengubah adapter saja
type userAdapter struct {
	repo userRepository
}

// FindUserById implements service.userAdapter.
func (u userAdapter) FindUserById(ctx context.Context, userId int) (user domain.User, err error) {
	userData, err := u.repo.FindUserById(ctx, userId)
	if err != nil {
		return
	}

	user = domain.User{
		Id:      userData.Id,
		Name:    userData.Name,
		Balance: userData.Balance,
	}

	return
}

// UpdateBalance implements service.userAdapter.
func (u userAdapter) UpdateBalance(ctx context.Context, userId int, balance int) (err error) {
	return u.repo.UpdateBalance(ctx, userId, balance)
}

func NewUserAdapter(repo userRepository) userAdapter {
	return userAdapter{
		repo: repo,
	}
}
