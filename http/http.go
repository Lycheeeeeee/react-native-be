package http

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	httptransport "github.com/go-kit/kit/transport/http"

	"github.com/minhkhiemm/example-go/endpoints"
)

// NewHTTPHandler ...
func NewHTTPHandler(endpoints endpoints.Endpoints,
	logger log.Logger,
	useCORS bool) http.Handler {
	r := chi.NewRouter()

	// if running on local (using `make dev`), include cors middleware
	if useCORS {
		cors := cors.New(cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			AllowCredentials: true,
		})
		r.Use(cors.Handler)
	}

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}

	r.Get("/_warm", httptransport.NewServer(
		endpoint.Nop,
		httptransport.NopRequestDecoder,
		httptransport.EncodeJSONResponse,
		options...,
	).ServeHTTP)

	r.Route("/orders", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.GetAllOrderByDate,
			DecodeGetAllByDateRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/", httptransport.NewServer(
			endpoints.CreateOrder,
			DecodeCreateOrderRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/byorderid/{orderid}", httptransport.NewServer(
			endpoints.GetOrderByID,
			DecodeGetOrderByIDRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Put("/{orderid}", httptransport.NewServer(
			endpoints.UpdateOrder,
			DecodeUpdateOrder,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/byshopid/{shop_id}", httptransport.NewServer(
			endpoints.GetOrderByShopID,
			DecodeGetORderByShopID,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Get("/bymonth/{month}", httptransport.NewServer(
			endpoints.GetOrderByMonth,
			DecodeGetOrderByMonth,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	r.Route("/accounts", func(r chi.Router) {
		r.Post("/", httptransport.NewServer(
			endpoints.CreateAccount,
			DecodeCreateAccount,
			encodeResponse,
			options...,
		).ServeHTTP)
		r.Post("/login", httptransport.NewServer(
			endpoints.Login,
			DecodeCreateAccount,
			encodeResponse,
			options...,
		).ServeHTTP)
	})
	r.Route("/drinks", func(r chi.Router) {
		r.Get("/", httptransport.NewServer(
			endpoints.GetAllDrink,
			DecodeNullRequest,
			encodeResponse,
			options...,
		).ServeHTTP)
	})

	return r
}

func DecodeNullRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return nil, nil
}
