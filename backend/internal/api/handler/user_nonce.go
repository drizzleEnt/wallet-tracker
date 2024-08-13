package handler

import (
	"net/http"

	"github.com/drizzleent/wallet-tracker/backend/converter"
	"github.com/drizzleent/wallet-tracker/backend/internal/api"
	"github.com/gin-gonic/gin"
)

const (
	idParams = "id"
)

func (h *handler) UserNonce(c *gin.Context) {
	id, ok := c.Params.Get(idParams)

	if !ok {
		api.NewErrorResponse(c, http.StatusBadRequest, "id is requared")
		return
	}

	if err := converter.ValidateId(id); err != nil {
		api.NewErrorResponse(c, http.StatusBadRequest, "invalid id"+err.Error())
		return
	}

	u, err := h.service.UserNonce(c, id)

	if err != nil {
		api.NewErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, u)

}
