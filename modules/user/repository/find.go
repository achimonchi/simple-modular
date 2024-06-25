package repository

import (
	"context"
	"errors"
	"simple-modular/internal/data"
	"simple-modular/modules/user/domain"
)

// FindAll implements service.repository.
func (r *repository) FindAll() (users []domain.User, err error) {
	var dbUsers = data.Users

	for _, user := range dbUsers {
		users = append(users, domain.User{
			Id:      user.Id,
			Name:    user.Name,
			Balance: user.Balance,
		})
	}

	return
}

// FindUserById implements adapter.userRepository.
func (r *repository) FindUserById(ctx context.Context, userId int) (user domain.User, err error) {
	var dbUsers = data.Users
	for _, u := range dbUsers {
		if u.Id == userId {
			user = domain.User{
				Id:      u.Id,
				Name:    u.Name,
				Balance: u.Balance,
			}
		}
	}

	if user.Id == 0 {
		err = errors.New("user not found")
		return
	}

	return
}
