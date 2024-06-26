package repository

import (
	"context"
	"simple-modular/internal/data"
)

// UpdateBalance implements adapter.userRepository.
func (r *repository) UpdateBalance(ctx context.Context, userId int, balance int) (err error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	var users = data.Users
	for i, user := range users {
		if user.Id == userId {
			users[i].Balance = balance
		}
	}

	return
}
