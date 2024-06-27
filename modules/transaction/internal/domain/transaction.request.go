package domain

type TransferRequest struct {
	From   int `json:"from"`
	To     int `json:"to"`
	Amount int `json:"amount"`
}

func (t TransferRequest) ToDomain() Transaction {
	return Transaction{
		Amount: t.Amount,
		From: User{
			Id: t.From,
		},
		To: User{
			Id: t.To,
		},
	}
}

type TopupRequest struct {
	From   int `json:"from"`
	Amount int `json:"amount"`
}

func (t TopupRequest) ToDomain() Transaction {
	return Transaction{
		Amount: t.Amount,
		From: User{
			Id: t.From,
		},
	}
}
