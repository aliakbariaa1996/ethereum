package v1

import (
	"github.com/aliakbariaa1996/ethereum/internal/services/ethereum"
	"github.com/labstack/echo/v4"
)

func Routes(router *echo.Echo, ether ethereum.UseService) {
	// Set up http handlers
	//var ctx *echo.Context
	hEther := NewEthereumHandler(ether)
	apiV1 := router.Group("/api/v1/ether")
	{
		apiV1.GET("/list", hEther.makeListTransaction)

	}
}
