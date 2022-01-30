package ethereum

type UseCase struct {
}

func NewEthereumUseCase() *UseCase {
	return &UseCase{}
}

type UseService interface {
	GetTransactions() (interface{}, error)
}
