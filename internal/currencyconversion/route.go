package currencyconversion

import (
	"context"
	"github.com/syrilster/go-microservice-example/internal/config"
	"net/http"
)

type ExchangeRateFetcher interface {
	FetchExchangeRate(ctx context.Context, req Request) (*Response, error)
}

func Route(rateFetcher ExchangeRateFetcher) (route config.Route) {
	route = config.Route{
		Path:    "/currency-converter/from/{from}/to/{to}/quantity/{quantity}",
		Method:  http.MethodGet,
		Handler: handler(rateFetcher),
	}

	return route
}
