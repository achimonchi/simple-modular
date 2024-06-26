package repository

import (
	"context"
	"simple-modular/internal/data"
	"simple-modular/modules/user/internal/domain"
)

// Insert implements service.repository.
func (r *repository) Insert(ctx context.Context, user domain.User) (err error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	var users = data.Users
	user.Id = len(users) + 1

	users = append(users, data.UserData{
		Id:      user.Id,
		Name:    user.Name,
		Balance: user.Balance,
	})

	data.Users = users

	return
}
