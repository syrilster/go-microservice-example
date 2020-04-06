package currencyexchange

type Request struct {
	FromCurrency string `json:"from"`
	ToCurrency   string `json:"to"`
}

type Response struct {
	FromCurrency       string  `json:"from"`
	ToCurrency         string  `json:"to"`
	ConversionMultiple float64 `json:"conversionMultiple"`
}
