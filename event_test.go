package myeventcase

import (
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const ConfirmDataStoreEventABI = "ConfirmDataStore(uint32,bytes32)"

var ConfirmDataStoreEventABIHash = crypto.Keccak256Hash([]byte(ConfirmDataStoreEventABI))

const DataLayrServiceManagerAddr = "0x5BD63a7ECc13b955C4F57e3F12A64c10263C14c1"

func TestEthClient_GetTxReceiptByHash(t *testing.T) {
	client, err := NewEthClient("https://rpc.mevblocker.io")
	if err != nil {
		t.Fatal(err)
	}
	receipt, err := client.GetTxReceiptByHash("0xfd26d40e17213bcafcf94bab9af92343302df9df970f20e1c9d515525e86e23e")
	if err != nil {
		t.Fatal(err)
	}
	abiUnt32, err := abi.NewType("uint32", "uint32", nil)
	if err != nil {
		t.Fatal(err)
	}
	abiBytes32, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		t.Fatal(err)
	}
	confirmDataStoreArgs := abi.Arguments{
		{
			Name:    "dataStoreId",
			Type:    abiUnt32,
			Indexed: false,
		},
		{
			Name:    "headerHash",
			Type:    abiBytes32,
			Indexed: false,
		},
	}
	dataStoreData := make(map[string]interface{})

	for _, log := range receipt.Logs {
		fmt.Println("address: ", log.Address.String())
		if !strings.EqualFold(log.Address.String(), DataLayrServiceManagerAddr) {
			continue
		}
		if log.Topics[0] != ConfirmDataStoreEventABIHash {
			continue
		}
		if len(log.Data) > 0 {
			err = confirmDataStoreArgs.UnpackIntoMap(dataStoreData, log.Data)
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println("dataStoreData.dataStoreId: ", dataStoreData["dataStoreId"])
			headerHash := dataStoreData["headerHash"].([32]byte)
			fmt.Println("dataStoreData.headerHash: ", common.Bytes2Hex(headerHash[:]))
		}
	}
}

func TestEthClient_GetLogs(t *testing.T) {
	client, err := NewEthClient("https://rpc.mevblocker.io")
	if err != nil {
		t.Fatal(err)
	}
	startblock := big.NewInt(20487721)
	endblock := big.NewInt(20487721)
	var contractAddressList []common.Address
	contractAddressList = append(contractAddressList, common.HexToAddress(DataLayrServiceManagerAddr))
	// var topics [][]common.Hash
	// topics = append(topics, []common.Hash{ConfirmDataStoreEventABIHash})
	logs, err := client.GetLogs(startblock, endblock, contractAddressList, nil)
	if err != nil {
		t.Fatal(err)
	}
	abiUnt32, err := abi.NewType("uint32", "uint32", nil)
	if err != nil {
		t.Fatal(err)
	}
	abiBytes32, err := abi.NewType("bytes32", "bytes32", nil)
	if err != nil {
		t.Fatal(err)
	}
	confirmDataStoreArgs := abi.Arguments{
		{
			Name:    "dataStoreId",
			Type:    abiUnt32,
			Indexed: false,
		},
		{
			Name:    "headerHash",
			Type:    abiBytes32,
			Indexed: false,
		},
	}
	dataStoreData := make(map[string]interface{})

	for _, log := range logs {
		fmt.Println("address: ", log.Address.String())
		if !strings.EqualFold(log.Address.String(), DataLayrServiceManagerAddr) {
			continue
		}
		if log.Topics[0] == ConfirmDataStoreEventABIHash {
			if len(log.Data) > 0 {
				err = confirmDataStoreArgs.UnpackIntoMap(dataStoreData, log.Data)
				if err != nil {
					t.Fatal(err)
				}
				fmt.Println("dataStoreData.dataStoreId: ", dataStoreData["dataStoreId"])
				headerHash := dataStoreData["headerHash"].([32]byte)
				fmt.Println("dataStoreData.headerHash: ", common.Bytes2Hex(headerHash[:]))
			}
		}
	}
}
