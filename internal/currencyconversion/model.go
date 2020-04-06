package currencyconversion

type Request struct {
	FromCurrency string  `json:"from"`
	ToCurrency   string  `json:"to"`
	Quantity     float64 `json:"quantity"`
}

type Response struct {
	Amount float64 `json:"amount"`
}
