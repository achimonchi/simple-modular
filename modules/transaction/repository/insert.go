package repository

import (
	"context"
	"simple-modular/internal/data"
	"simple-modular/modules/transaction/domain"
	"time"
)

// Insert implements service.repository.
func (r *repository) Insert(ctx context.Context, req domain.Transaction) (err error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	trxs := data.Transactions
	req.Id = len(trxs) + 1

	trxs = append(trxs, data.TransactionData{
		Id:        req.Id,
		Amount:    req.Amount,
		From:      req.From.Id,
		To:        req.To.Id,
		Type:      "transfer",
		CreatedAt: time.Now(),
	})

	data.Transactions = trxs

	return
}

// Topup implements service.repository.
func (r *repository) Topup(ctx context.Context, req domain.Transaction) (err error) {
	r.mut.Lock()
	defer r.mut.Unlock()

	trxs := data.Transactions
	req.Id = len(trxs) + 1

	trxs = append(trxs, data.TransactionData{
		Id:        req.Id,
		Amount:    req.Amount,
		From:      req.From.Id,
		Type:      "topup",
		CreatedAt: time.Now(),
	})

	data.Transactions = trxs

	return
}
