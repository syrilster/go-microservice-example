package currencyconversion

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Request struct {
	FromCurrency string  `json:"from"`
	ToCurrency   string  `json:"to"`
	Quantity     float64 `json:"quantity"`
}

type Response struct {
	Amount float64 `json:"amount"`
}

func (request Request) Validate() error {
	return validation.ValidateStruct(&request,
		validation.Field(&request.FromCurrency, validation.Required),
		validation.Field(&request.ToCurrency, validation.Required),
		validation.Field(&request.Quantity, validation.Required),
	)
}
