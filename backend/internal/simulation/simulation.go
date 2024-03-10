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

func Simulation(wsClient *ethclient.Client, richPrivKey *ecdsa.PrivateKey, richPubKey common.Address, numWallets int, ethersPerWallets int, ethersPerTransactions float64, numTransactions int, stopChan chan bool) {
	walletsFrom = wallets.CreateWallets(numWallets)
	faucet.SendEthers(wsClient, richPrivKey, richPubKey, walletsFrom, float64(ethersPerWallets))

	walletsTo = wallets.CreateWallets(numWallets)
	faucet.SendEthers(wsClient, richPrivKey, richPubKey, walletsTo, float64(ethersPerWallets))

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
			fmt.Println("Simulation : Going for block number:", (block.Number().Uint64() + 1))
			faucet.SendEthersFromAPoolToAPool(wsClient, walletsFrom, walletsTo, numTransactions, ethersPerTransactions)
		case <-stopChan:
			fmt.Println("Simulation : Stopping simulation...")
			return
		}
	}

}
