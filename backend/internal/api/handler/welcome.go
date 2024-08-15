package handler

import (
	"net/http"

	"github.com/drizzleent/wallet-tracker/backend/internal/api"
	"github.com/drizzleent/wallet-tracker/backend/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) Welcome(c *gin.Context) {
	value, exist := c.Get("user")
	if !exist {
		api.NewErrorResponse(c, http.StatusUnauthorized, "user dont auththorized")
		return
	}

	user := value.(*model.User)

	resp := struct {
		Msg string `json:""msg`
	}{
		Msg: user.Address,
	}

	c.JSON(http.StatusOK, resp)
}
