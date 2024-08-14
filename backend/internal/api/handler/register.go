package handler

import (
	"net/http"

	"github.com/drizzleent/wallet-tracker/backend/internal/api"
	"github.com/drizzleent/wallet-tracker/backend/internal/converter"
	"github.com/drizzleent/wallet-tracker/backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) Register(c *gin.Context) {
	var p model.RegisterPayload
	if err := c.Bind(&p); err != nil {
		api.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := converter.ValidateAddress(p.Address); err != nil {
		api.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	u, err := h.service.Register(c, &p)

	if err != nil {
		//TODO: errors
		api.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, u)

}
