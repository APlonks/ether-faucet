package faucet

import (
	"backend/internal/utils"
	"backend/internal/wallets"
	"context"
	"crypto/ecdsa"
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
func SendTransactionLegacy(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, toWallet common.Address, nbEthers float64) {
	var (
		nonce uint64
		err   error
	)
	nonce, err = client.PendingNonceAt(context.Background(), fromAddress)
	utils.ErrManagement(err)
	amount := big.NewFloat(nbEthers)

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
	tx := types.NewTransaction(nonce, toWallet, value, gasLimit, gasPrice, data)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal("Problem signing transaction:", err)
	}
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal("Problem sending transaction:", err)
	}
}

func SendEthersToSpecificWallet(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, toWallet wallets.Wallet, nbEthers float64) {

	SendTransactionLegacy(client, privateKey, fromAddress, toWallet.Address, nbEthers)
}

func SendEthers(client *ethclient.Client, privateKey *ecdsa.PrivateKey, fromAddress common.Address, wallets []wallets.Wallet, nbEthers float64) {

	for _, wallet := range wallets {
		SendTransactionLegacy(client, privateKey, fromAddress, wallet.Address, nbEthers)
	}
}
