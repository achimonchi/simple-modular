package service

import (
	"context"
	"errors"
	"log"
	"simple-modular/modules/transaction/internal/domain"
)

func (s service) Transfer(ctx context.Context, req domain.Transaction) (err error) {
	if s.user == nil {
		return errors.New("user adapter is nil")
	}

	// get current balance
	fromUser, err := s.user.FindUserById(ctx, req.From.Id)
	if err != nil {
		log.Println("[Transfer, FindUserById] error : ", err)
		return err
	}
	toUser, err := s.user.FindUserById(ctx, req.To.Id)
	if err != nil {
		log.Println("[Transfer, FindUserById] error : ", err)
		return err
	}

	req.SetFromUser(fromUser)
	req.SetToUser(toUser)

	if err := req.Validate(); err != nil {
		log.Println("[Transfer, ValidateBalance] error : ", err)
		return err
	}

	if err = s.repo.Insert(ctx, req); err != nil {
		log.Println("[Transfer, Insert] error : ", err)
		return err
	}

	newFromUser := req.GetNewFromUserData()
	newToUser := req.GetNewToUserData()

	if err = s.user.UpdateBalance(ctx, newFromUser.Id, newFromUser.Balance); err != nil {
		log.Println("[Transfer, UpdateBalance From User] error : ", err)
		return err
	}

	if err = s.user.UpdateBalance(ctx, newToUser.Id, newToUser.Balance); err != nil {
		log.Println("[Transfer, UpdateBalance To User] error : ", err)
		return err
	}

	return
}

func (s service) Topup(ctx context.Context, req domain.Transaction) (err error) {
	if s.user == nil {
		return errors.New("user adapter is nil")
	}

	// get current balance
	fromUser, err := s.user.FindUserById(ctx, req.From.Id)
	if err != nil {
		log.Println("[Topup, FindUserById] error : ", err)
		return err
	}
	req.SetFromUser(fromUser)

	if err = s.repo.Topup(ctx, req); err != nil {
		log.Println("[Topup, Insert] error : ", err)
		return err
	}

	updatedUser := req.GetAfterTopup()

	if err = s.user.UpdateBalance(ctx, updatedUser.Id, updatedUser.Balance); err != nil {
		log.Println("[Topup, UpdateBalance From User] error : ", err)
		return err
	}

	return
}
