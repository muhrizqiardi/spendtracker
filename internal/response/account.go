package response

type CommonAccountResponse struct {
	ID            uint   `json:"id"`
	UserID        uint   `json:"userId"`
	CurrencyID    uint   `json:"currencyId"`
	Name          string `json:"name"`
	InitialAmount int    `json:"initialAmount"`
}
