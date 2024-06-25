package domain

import (
	"errors"
	"time"
)

type Transaction struct {
	Id        int
	From      User
	To        User
	Amount    int
	Type      string
	CreatedAt time.Time
}

func (t *Transaction) SetFromUser(user User) {
	t.From = user
}

func (t *Transaction) SetToUser(user User) {
	t.To = user
}

func (t *Transaction) Validate() error {
	if err := t.ValidateUser(); err != nil {
		return err
	}

	if err := t.ValidateBalance(); err != nil {
		return err
	}

	return nil
}

func (t *Transaction) ValidateUser() error {
	if t.From.Id == t.To.Id {
		return errors.New("user from and target is same")
	}

	return nil
}

func (t *Transaction) ValidateBalance() error {
	if t.From.Balance < t.Amount {
		return errors.New("insufficient balance")
	}
	return nil
}

func (t *Transaction) GetNewFromUserData() User {
	return User{
		Id:      t.From.Id,
		Name:    t.From.Name,
		Balance: t.From.Balance - t.Amount,
	}
}
func (t *Transaction) GetAfterTopup() User {
	return User{
		Id:      t.From.Id,
		Name:    t.From.Name,
		Balance: t.From.Balance + t.Amount,
	}
}

func (t *Transaction) GetNewToUserData() User {
	return User{
		Id:      t.To.Id,
		Name:    t.To.Name,
		Balance: t.To.Balance + t.Amount,
	}
}
