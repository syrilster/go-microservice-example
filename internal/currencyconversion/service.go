package currencyconversion

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/syrilster/go-microservice-example/internal/currencyexchange"
	"math"
	"strconv"
)

type Service struct {
	client currencyexchange.ClientInterface
}

func NewService(c currencyexchange.ClientInterface) *Service {
	return &Service{c}
}

func (service Service) FetchExchangeRate(ctx context.Context, request Request) (*Response, error) {
	ctxLogger := log.WithContext(ctx)
	ctxLogger.Infof("Calling the currency exchange service fetch API")

	req := toCurrencyExchangeRequest(request)
	currencyExchangeResp, err := service.client.GetExchangeRate(ctx, req)
	if err != nil {
		ctxLogger.Infof("Failed to fetch currency exchange rate: %v", err)
		return nil, err
	}

	conversionMultiple, err := strconv.ParseFloat(currencyExchangeResp.ConversionMultiple, 64)
	if err != nil {
		ctxLogger.Infof("Failed to un marshall the exchange rate: %v", err)
		return nil, err
	}

	response := &Response{Amount: calculateAmount(request.Quantity, conversionMultiple)}
	return response, nil
}

func calculateAmount(quantity float64, conversionMultiple float64) float64 {
	amount := conversionMultiple * quantity
	return math.Round(amount*100) / 100
}

func toCurrencyExchangeRequest(request Request) currencyexchange.Request {
	return currencyexchange.Request{
		FromCurrency: request.FromCurrency,
		ToCurrency:   request.ToCurrency,
	}
}
