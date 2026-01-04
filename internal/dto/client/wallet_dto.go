package client

type DepositWalletDTO struct {
	Amount float64 `json:"amount" binding:"required,min=1"`
}
