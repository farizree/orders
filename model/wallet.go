package model

type Wallet struct {
	ID         string  `json:"id"`
	UserID     string  `json:"user_id"`
	Balance    float64 `json:"balance"`
	Status     int64   `json:"status"`
	CreatedAt  int64   `json:"created_at"`
	UpdatedAt  int64   `json:"updated_at"`
	IsActive   int64   `json:"is_active"`
	CurrencyID int64   `json:"currency_id"`
}

type WalletByUserID struct {
	UserID string `form:"user_id"`
}

type RequestListWallet struct {
	UserID string

	Page int32
	Rows int32
}

type ResponseWallet struct {
	*Wallet
}

type ResponseListWallet struct {
	Wallets []*Wallet `json:"wallets"`
	Total   int32     `json:"total"`
}
