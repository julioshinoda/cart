package handler

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/julioshinoda/cart/entity"
	"github.com/julioshinoda/cart/internal/api/presenter"
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
	r.Post("/{id}/item", c.AddItem)
	r.Put("/{id}/item", c.UpdateItem)
	r.Delete("/{id}/item/{item-id}", c.DeleteItem)
	r.Delete("/{id}/clean", c.ClearCart)
	r.Post("/{id}/coupon", c.AddCoupon)
}

func (c *CartHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	crt, err := c.service.GetCart(chi.URLParam(r, "id"))

	if errors.Is(err, entity.ErrNotFound) {
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusNotFound, response)
		return
	}

	if err != nil {
		c.logger.Error(
			"Error on get cart info",
			zap.String("id", chi.URLParam(r, "id")),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	response, _ := json.Marshal(crt)
	presenter.MountResponse(w, http.StatusOK, response)
}

func (c *CartHandler) AddItem(w http.ResponseWriter, r *http.Request) {
	var item entity.Items
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		c.logger.Error(
			"Error on decode cart item",
			zap.String("id", chi.URLParam(r, "id")),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	_, err := c.service.AddItem(chi.URLParam(r, "id"), item)
	if err != nil {
		c.logger.Error(
			"Error on add cart item",
			zap.String("id", chi.URLParam(r, "id")),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	presenter.MountResponse(w, http.StatusOK, nil)
}

func (c *CartHandler) DeleteItem(w http.ResponseWriter, r *http.Request) {
	if err := c.service.RemoveItem(
		chi.URLParam(r, "id"),
		chi.URLParam(r, "item-id"),
	); err != nil {
		c.logger.Error(
			"Error on remove cart item",
			zap.String("id", chi.URLParam(r, "id")),
			zap.String("item-id", chi.URLParam(r, "item-id")),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	presenter.MountResponse(w, http.StatusOK, nil)
}

func (c *CartHandler) ClearCart(w http.ResponseWriter, r *http.Request) {
	if err := c.service.ClearCart(chi.URLParam(r, "id")); err != nil {
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	presenter.MountResponse(w, http.StatusOK, nil)
}

func (c *CartHandler) UpdateItem(w http.ResponseWriter, r *http.Request) {
	var item entity.Items
	if err := json.NewDecoder(r.Body).Decode(&item); err != nil {
		c.logger.Error(
			"Error on decode cart item",
			zap.String("id", chi.URLParam(r, "id")),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	if err := c.service.UpdateCart(chi.URLParam(r, "id"), item); err != nil {
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	presenter.MountResponse(w, http.StatusOK, nil)
}

func (c *CartHandler) AddCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon entity.Coupon
	if err := json.NewDecoder(r.Body).Decode(&coupon); err != nil {
		c.logger.Error(
			"Error on decode coupon",
			zap.String("id", chi.URLParam(r, "id")),
			zap.Error(err),
		)
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	if err := c.service.AddCoupon(chi.URLParam(r, "id"), coupon); err != nil {
		response, _ := json.Marshal(presenter.ErrorMessage{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		presenter.MountResponse(w, http.StatusBadRequest, response)
		return
	}
	presenter.MountResponse(w, http.StatusOK, nil)
}
