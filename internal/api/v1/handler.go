package v1

import "github.com/aliakbariaa1996/ethereum/internal/services/ethereum"

type EthereumHandler struct {
	ethereumService ethereum.UseService

}
func NewEthereumHandler(etherSer ethereum.UseService) *EthereumHandler {
	return &EthereumHandler{
		ethereumService: etherSer,
	}
}