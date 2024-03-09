package main

import (
	"backend/internal/faucet"
	"backend/internal/simulation"
	"backend/internal/utils"
	"backend/internal/wallets"
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	configPath  string
	err         error
	config      *utils.Config
	clientHttp  *ethclient.Client
	clientWs    *ethclient.Client
	richPrivKey *ecdsa.PrivateKey
	richPubKey  common.Address
	stopChannel chan bool // Simulation control
)

func main() {

	configPath, err = utils.ParseFlags()
	config, err = utils.LoadConfig(configPath)

	router := gin.Default()

	router.Use(cors.Default()) // Allow all

	router.POST("/faucet", SendEthersToSpecificAddress)

	router.POST("/start-simulation", StartSimulationHandler)

	router.POST("/stop-simulation", StopSimulationHandler)

	router.Run() // listen and serve on 0.0.0.0:8080
}

func SendEthersToSpecificAddress(c *gin.Context) {

	type UserRequest struct {
		ToWallet string `json:"wallet"`
	}
	var userReq UserRequest

	// Config Client //////////////
	richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(config.Connection.Rich_private_key)
	utils.ErrManagement(err)

	clientHttp, err = ethclient.Dial(config.Connection.Http_endpoint)
	utils.ErrManagement(err)
	//////////////////////////////

	if err := c.BindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userReq.ToWallet == "" {
		fmt.Println("Send 1 ether to the Specific Address : 0x0000000000000000000000000000000000000000")
		faucet.SendEthersToSpecificAddress(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), config.Simulation.Ethers)
		c.JSON(200, gin.H{"status": "Request sent to the backend"})
	} else {
		valid := utils.IsValidAddress("userReq.ToWallet")
		if valid {
			fmt.Println("Send 1 ether to the Specific Address :", userReq.ToWallet)
			faucet.SendEthersToSpecificAddress(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), config.Simulation.Ethers)
			c.JSON(200, gin.H{"status": "Request sent to the backend"})
		} else {
			c.JSON(200, gin.H{"status": "Public address format is not valid"})
		}

	}

}

func StartSimulationHandler(c *gin.Context) {
	go Simulation(c, stopChannel)
	c.JSON(http.StatusOK, gin.H{"message": "Simulation started"})
}

func StopSimulationHandler(c *gin.Context) {
	// Send stop signal
	stopChannel <- true
	c.JSON(http.StatusOK, gin.H{"message": "Simulation stopping signal sent"})
}

func Simulation(c *gin.Context, stopChan chan bool) {

	type UserRequest struct {
		AccountsPerWallets    int `json:"accounts_per_wallets"`
		EthersPerTransactions int `json:"ethers_per_transactions"`
		TransactionsPerBlocks int `json:"transactions_per_blocks"`
	}
	var userReq UserRequest

	select {
	case <-stopChan:
		fmt.Println("Stopping simulation...")
		return
	case <-time.After(10 * time.Second):
		// Config Client //////////////
		richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(config.Connection.Rich_private_key)
		utils.ErrManagement(err)

		clientWs, err = ethclient.Dial(config.Connection.Ws_endpoint)
		utils.ErrManagement(err)
		//////////////////////////////

		if err := c.BindJSON(&userReq); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		simulation.Simulation(clientWs, richPrivKey, richPubKey, userReq.AccountsPerWallets, userReq.EthersPerTransactions, userReq.TransactionsPerBlocks)
	}

}
