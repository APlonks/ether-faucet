package simulation

import (
	"backend/internal/faucet"
	"backend/internal/wallets"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	walletsFrom []wallets.Wallet
	walletsTo   []wallets.Wallet
)

func Simulation(wsClient *ethclient.Client, richPrivKey *ecdsa.PrivateKey, richPubKey common.Address, numWallets int, nbEthers int, numTransactions int) {
	walletsFrom = wallets.CreateWallets(numWallets)
	fmt.Println("First list of accounts")
	for _, wallet := range walletsFrom {
		fmt.Println("Public key:", wallet.AddressHex, "; Private key:", wallet.KeyHex)
	}
	faucet.SendEthers(wsClient, richPrivKey, richPubKey, walletsFrom, 1)

	walletsTo = wallets.CreateWallets(numWallets)
	fmt.Println("Second list of accounts")
	for _, wallet := range walletsTo {
		fmt.Println("Public key:", wallet.AddressHex, "; Private key:", wallet.KeyHex)
	}
	faucet.SendEthers(wsClient, richPrivKey, richPubKey, walletsTo, 1)

	time.Sleep(13 * time.Second) // Waiting for a block

	headers := make(chan *types.Header)
	sub, err := wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			block, err := wsClient.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println("Going for block number:", (block.Number().Uint64() + 1)) // 3477413
			fmt.Println(wsClient.BalanceAt(context.Background(), walletsFrom[0].Address, nil))
			faucet.SendEthersFromAPoolToAPool(wsClient, walletsFrom, walletsTo, numTransactions)
			fmt.Println(numTransactions, "transactions sended")
		}
	}

}
