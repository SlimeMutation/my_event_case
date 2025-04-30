package myeventcase

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
)

type EthClient struct {
	client *ethclient.Client
}

func NewEthClient(rpcUrl string) (*EthClient, error) {
	client, err := ethclient.DialContext(context.Background(), rpcUrl)
	if err != nil {
		log.Error("NewEthClient fail", "err", err)
		return nil, err
	}
	return &EthClient{client: client}, nil
}

func (e *EthClient) GetTxReceiptByHash(txHash string) (*types.Receipt, error) {
	return e.client.TransactionReceipt(context.Background(), common.HexToHash(txHash))
}

func (e *EthClient) GetLogs(startBlock, endBlock *big.Int, contractAddressList []common.Address, topics [][]common.Hash) ([]types.Log, error) {
	filterQueryParams := ethereum.FilterQuery{
		FromBlock: startBlock,
		ToBlock:   endBlock,
		Addresses: contractAddressList,
		Topics:    topics,
	}
	return e.client.FilterLogs(context.Background(), filterQueryParams)
}
