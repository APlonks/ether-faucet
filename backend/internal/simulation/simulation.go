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

func Simulation(wsClient *ethclient.Client, richPrivKey *ecdsa.PrivateKey, richPubKey common.Address, numWallets int, ethersPerWallet int, ethersPerTransaction float64, numTransactions int, stopChan chan bool) error {
	walletsFrom = wallets.CreateWallets(numWallets)
	faucet.SendEthersToAWalletPool(wsClient, richPrivKey, richPubKey, walletsFrom, float64(ethersPerWallet))

	walletsTo = wallets.CreateWallets(numWallets)
	faucet.SendEthersToAWalletPool(wsClient, richPrivKey, richPubKey, walletsTo, float64(ethersPerWallet))

	time.Sleep(13 * time.Second) // Waiting for a block

	headers := make(chan *types.Header)
	sub, err := wsClient.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		fmt.Println("Error while trying subcribbe to the blockchain")
		return err
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
			faucet.SendEthersFromAPoolToAPool(wsClient, walletsFrom, walletsTo, numTransactions, ethersPerTransaction)
		case <-stopChan:
			fmt.Println("Simulation : Stopping simulation...")
			return nil
		}
	}

}
