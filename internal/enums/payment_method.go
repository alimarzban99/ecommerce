package enums

type PaymentMethod string

const (
	PaymentGateway PaymentMethod = "gateway" // پرداخت آنلاین
	PaymentWallet  PaymentMethod = "wallet"  // پرداخت از کیف پول
)
