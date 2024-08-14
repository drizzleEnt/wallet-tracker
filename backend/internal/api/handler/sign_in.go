package handler

import (
	"net/http"

	"github.com/drizzleent/wallet-tracker/backend/internal/api"
	"github.com/drizzleent/wallet-tracker/backend/internal/converter"
	"github.com/drizzleent/wallet-tracker/backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) Signin(c *gin.Context) {
	var p model.SigningPayload

	if err := c.Bind(&p); err != nil {
		api.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := converter.ValidateSignPayload(&p); err != nil {
		api.NewErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.service.Signin(&p)

	if err != nil {
		api.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resp)

}
