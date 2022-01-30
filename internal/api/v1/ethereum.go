package v1

import (
	httpx "github.com/aliakbariaa1996/ethereum/internal/http"
	"github.com/labstack/echo/v4"
	"net/http"
)

// @Summary makeListTransaction
// @Tags transaction
// @Description get transactions based on project address.
// @ID makeListTransaction
// @Accept  json
// @Produce  json
// @Param input body requests.MakeListTransaction true "credentials"
// @Success 200 {object} transactionRes
// @Failure 400 {object} errorx.Error
// @Router api/v1/ether/ether [get]
type transactionRes struct {
	List interface{}
}

func (h *EthereumHandler) makeListTransaction(c echo.Context) error {
	res, err := h.ethereumService.GetTransactions()
	if err != nil {
		httpx.JSONResponse(c.Response(), err, http.StatusBadRequest)
		return err
	}
	resp := transactionRes{List: res}
	httpx.JSONResponse(c.Response(), resp, http.StatusOK)
	return err
}
