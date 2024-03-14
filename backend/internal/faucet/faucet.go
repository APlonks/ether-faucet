package faucet

import (
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

func SendEthersFromAPoolToAPool(client *ethclient.Client, walletsFrom []wallets.Wallet, walletsTo []wallets.Wallet, numTransactions int, ethersPerTransaction float64) {

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < numTransactions/2; i++ {
			indexFrom := rand.IntN(cap(walletsFrom))
			indexTo := rand.IntN(cap(walletsTo))
			SendEthersToSpecificWallet(client, &walletsFrom[indexFrom].Key, walletsFrom[indexFrom].Address, walletsTo[indexTo], ethersPerTransaction)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	func() {
		for i := 0; i < numTransactions/2; i++ {
			indexFrom := rand.IntN(cap(walletsFrom))
			indexTo := rand.IntN(cap(walletsTo))
			SendEthersToSpecificWallet(client, &walletsTo[indexTo].Key, walletsTo[indexTo].Address, walletsFrom[indexFrom], ethersPerTransaction)
			time.Sleep(time.Millisecond * 10)
		}
	}()
	wg.Wait()
}

// Transaction Pre EIP 1559
func SendTransactionLegacy(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, toWallet common.Address, nbEthers float64) error {
	var (
		nonce uint64
		err   error
	)
	// fmt.Println("The client:", client)
	// fmt.Println("The privateKey:", privateKey)
	// fmt.Println("The fromAddres:", fromAddress)
	// fmt.Println("The toWallet:", toWallet)
	// fmt.Println("The number of ethers:", nbEthers)
	fmt.Println("The fromAddress:", fromAddress)
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	fmt.Println("The nonce:", nonce)
	if err != nil {
		fmt.Println("Error while trying to retrieve the Nonce:", err)
		return err
	}
	amount := big.NewFloat(nbEthers)
	fmt.Println("PendingNonce At passed")

	// Convert Ethers to Wei (1 Ether = 1e18 Wei)
	weiValue := new(big.Float).Mul(amount, big.NewFloat(1e18))
	value := new(big.Int)

	weiValue.Int(value)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	var data []byte
	fmt.Println("Pre NewTransaction")
	tx := types.NewTransaction(nonce, toWallet, value, gasLimit, gasPrice, data)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Pre signing transaction")
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("Problem signing transaction:", err)
	}
	fmt.Println("Post signing transaction")
	fmt.Println("Pre sending transaction")
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Problem sending transaction:", err)
	}
	fmt.Println("Post sending transaction")

	return nil
}

func SendEthersToSpecificWallet(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, toWallet wallets.Wallet, nbEthers float64) {

	SendTransactionLegacy(client, privateKey, fromAddress, toWallet.Address, nbEthers)
}

func SendEthersToAWalletPool(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, wallets []wallets.Wallet, nbEthers float64) {

	for _, wallet := range wallets {
		SendTransactionLegacy(client, privateKey, fromAddress, wallet.Address, nbEthers)
	}
}
