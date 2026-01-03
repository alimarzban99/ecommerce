package enums

type TransactionType string

const (
	TransactionDeposit TransactionType = "deposit"
	TransactionRefund  TransactionType = "refund"
	TransactionPayment TransactionType = "payment"
)
