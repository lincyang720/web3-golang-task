package home1

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func connectToEthereum() (*ethclient.Client, error) {
	client, err := ethclient.Dial("https://ethereum-sepolia-rpc.publicnode.com")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	return client, err
}

func BlockQuery() {
	client, err := connectToEthereum()
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(5671744)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("区块号:", block.Number().Uint64())
	fmt.Println("区块哈希:", block.Hash().Hex())
	fmt.Println("时间戳:", block.Time())
	fmt.Println("交易数量:", len(block.Transactions()))
}

func SendTransaction() {
	client, err := connectToEthereum()
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("5fb0d527911392bd021a6cad58c90bc5794b1a7354f5840f38170c758f95c6a7")
	if err != nil {
		log.Fatal(err)
	}

	// 获取公钥和地址
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// 获取账户地址和nonce
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// 设置交易参数
	value := big.NewInt(1000000000)
	gasLimit := uint64(21000)
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x524d25aF35803B4F4E4Ca6232a82AD3270891159")
	var data []byte
	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("交易已发送: %s\n", signedTx.Hash().Hex())

}
