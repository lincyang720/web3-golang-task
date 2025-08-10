package home2

import (
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/lincyang720/web3-golang-task/dapp/task1/counter"
)

const (
	contractAddress = "0xd6A42bF961B4A7360741240B28E592A3ba9876f4"
)

func GenerateContract() {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	if err != nil {
		log.Fatal(err)
	}

	conterContract, err := counter.NewCounter(common.HexToAddress(contractAddress), client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		log.Fatal(err)
	}

	opt, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
	if err != nil {
		log.Fatal(err)
	}

	tx, err := conterContract.Increment(opt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("tx hash:", tx.Hash().Hex())
	fmt.Println("tx value:", tx.Value().String())
}
