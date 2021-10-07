package handler

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/julioshinoda/cart/internal/usecase/cart"
	"go.uber.org/zap"
)

type CartHandler struct {
	logger  *zap.Logger
	service cart.UseCase
}

func NewCartHandlers(logger *zap.Logger, service cart.UseCase) *CartHandler {
	return &CartHandler{
		logger:  logger,
		service: service,
	}
}

func MakeCartHandlers(r chi.Router, c *CartHandler) {

	r.Get("/{id}", c.GetByID)

}

func (c *CartHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	crt, err := c.service.GetCart(chi.URLParam(r, "id"))
	if err != nil {
		c.logger.Error(
			"Error on get cart info",
			zap.String("id", chi.URLParam(r, "id")),
			zap.Error(err),
		)
	}
	response, _ := json.Marshal(crt)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
