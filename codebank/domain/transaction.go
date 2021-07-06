package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TransactionRepository interface {
	SaveTransaction(Transaction Transaction, creditCard CreditCard) error
	GetCreditCard(creditCard CreditCard) (CreditCard, error)
	CreateCreditCard(creditCard CreditCard) error
}

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Store        string
	Description  string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	t := &Transaction{}
	t.ID = uuid.NewV4().String()
	t.CreatedAt = time.Now()
	return t
}

func (t *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if t.Amount+creditCard.Balance > creditCard.Limit {
		t.Status = "rejected"
		return
	}

	t.Status = "approved"
	creditCard.Balance = creditCard.Balance + t.Amount
}
