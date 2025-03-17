package entities

import (
	"gorm.io/gorm"
)

// CryptomusTransaction is a struct that represents a transaction
// by Cryptomus.
type CryptomusTransaction struct {
	gorm.Model

	// TransactionID is a unique identifier for the transaction
	// associated with this Cryptomus transaction.
	TransactionID uint `json:"transaction_id"`

	// UUID indicates Invoice uuid.
	UUID string `json:"uuid"`
	// OrderID indicates Order ID in the website.
	OrderID string `json:"order_id"`
	// Amount stores the amount of the invoice.
	Amount string `json:"amount"`
	// PaymentAmount stores the amount paid by client.
	PaymentAmount *string `json:"payment_amount"`
	// PaymentAmountUSD stores the amount paid by client in USD.
	PaymentAmountUSD *string `json:"payment_amount_usd"`
	// IsFinal indicates whether the invoice is finalized.
	// When invoice is finalized, it is impossible to pay an invoice
	// (it's either paid or expired).
	IsFinal bool `json:"is_final"`
	// PaymentStatus indicates at what stage the payment is at the moment.
	PaymentStatus CryptomusPaymentStatus `json:"payment_status"`
	// From indicates the wallet address from which the payment was made.
	From *string `json:"paid_from"`
	// Network indicates blockchain network code.
	Network *string `json:"network"`
	// Currency indicates invoice currency code.
	Currency string `json:"currency"`
	// TxID stores transaction hash.
	TxID *string `json:"txid"`
}

// TableName overrides the default table name
func (CryptomusTransaction) TableName() string {
	return "tbl_cryptomus_transactions"
}
