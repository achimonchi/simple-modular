package service

import (
	"context"
	"errors"
	"log"
)

// UpdateBalance implements domainhandler.service.
func (s service) UpdateBalance(ctx context.Context, userId int, balance int) (err error) {
	user, err := s.repo.FindUserById(ctx, userId)
	if err != nil {
		log.Println("[UpdateBalance, FindUserById] error :", err.Error())
		return
	}

	if !user.IsExists() {
		err = errors.New("user not found")
		log.Println("[UpdateBalance, IsExists] error :", err.Error())
		return
	}

	if err = s.repo.UpdateBalance(ctx, userId, balance); err != nil {
		log.Println("[UpdateBalance, UpdateBalance] error :", err.Error())
		return
	}

	return
}
