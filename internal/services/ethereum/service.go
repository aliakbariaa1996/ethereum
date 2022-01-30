package ethereum

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func (u *UseCase) GetTransactions() (interface{}, error) {
	var transactionList []string
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/5e0125260a82488c893a5b59fa12c74c")
	if err != nil {
		log.Fatal(err)
	}
	blockHash := common.HexToHash("0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		transactionList = append(transactionList, tx.Hash().Hex())
	}
	return transactionList, err
}
