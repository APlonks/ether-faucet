package faucet

import (
	"backend/internal/utils"
	"backend/internal/wallets"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	wg sync.WaitGroup
)

func SendEthersFromAPoolToAPool(client *ethclient.Client, walletsFrom []wallets.Wallet, walletsTo []wallets.Wallet, numTransactions int) {

	var (
		err      error
		nbEthers int
	)
	_ = err
	nbEthers = 1
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numTransactions/2; i++ {
			indexFrom := rand.IntN(cap(walletsFrom))
			indexTo := rand.IntN(cap(walletsTo))
			SendEthersToSpecificWallet(client, &walletsFrom[indexFrom].Key, walletsFrom[indexFrom].Address, walletsTo[indexTo], nbEthers)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	func() {
		for i := 0; i < numTransactions/2; i++ {
			indexFrom := rand.IntN(cap(walletsFrom))
			indexTo := rand.IntN(cap(walletsTo))
			SendEthersToSpecificWallet(client, &walletsTo[indexTo].Key, walletsTo[indexTo].Address, walletsFrom[indexFrom], nbEthers)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	wg.Wait()
}

func SendEthersToSpecificAddress(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, toWallet common.Address, nbEthers int) {
	var (
		nonce uint64
		err   error
	)
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	utils.ErrManagement(err)
	// Convert nbEthers (int) en big.Int
	amount := big.NewInt(int64(nbEthers))
	// Convert Ethers to Wei (1 Ether = 1e18 Wei)
	weiValue := new(big.Int).Mul(amount, big.NewInt(1000000000000000000))

	value := weiValue         // in wei (1 eth)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	tx := types.NewTransaction(nonce, toWallet, value, gasLimit, gasPrice, data)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("LA", err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("ICI:", err)
	}
	fmt.Println("Transaction")
}

func SendEthersToSpecificWallet(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, toWallet wallets.Wallet, nbEthers int) {
	var (
		nonce uint64
		err   error
	)
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	utils.ErrManagement(err)
	// Convert nbEthers (int) en big.Int
	amount := big.NewInt(int64(nbEthers))
	// Convert Ethers to Wei (1 Ether = 1e18 Wei)
	weiValue := new(big.Int).Mul(amount, big.NewInt(0))

	value := weiValue         // in wei (1 eth)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	tx := types.NewTransaction(nonce, toWallet.Address, value, gasLimit, gasPrice, data)
	chainID, err := client.ChainID(context.Background())
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
	fmt.Println("Transaction")
}

func SendEthers(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, wallets []wallets.Wallet, nbEthers int) {

	var (
		nonce uint64
		err   error
	)

	for _, wallet := range wallets {
		nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
		utils.ErrManagement(err)
		fmt.Println("Nonce:", nonce)
		// Convert nbEthers (int) en big.Int
		amount := big.NewInt(int64(nbEthers))
		// Convert Ethers to Wei (1 Ether = 1e18 Wei)
		weiValue := new(big.Int).Mul(amount, big.NewInt(1e18))

		value := weiValue         // in wei (1 eth)
		gasLimit := uint64(21000) // in units
		gasPrice, err := client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		var data []byte
		tx := types.NewTransaction(nonce, wallet.Address, value, gasLimit, gasPrice, data)
		chainID, err := client.ChainID(context.Background())
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
		// fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
		// fmt.Println()
	}
}
