package main

import (
	"backend/internal/faucet"
	"backend/internal/simulation"
	"backend/internal/utils"
	"backend/internal/wallets"
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"reflect"

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
	SimuRunning bool
)

type SimuRequest struct {
	AccountsPerWallet    int     `json:"accounts_per_wallet"`
	EthersPerWallet      int     `json:"ethers_per_wallet"`
	EthersPerTransaction float64 `json:"ethers_per_transaction"`
	TransactionsPerBlock int     `json:"transactions_per_block"`
}

func main() {

	// Init channel stopChannel
	stopChannel = make(chan bool)

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
		fmt.Println("Faucet : Send 1 ether to the Specific Address : 0x0000000000000000000000000000000000000000")
		faucet.SendTransactionLegacy(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), float64(1))
		c.JSON(200, gin.H{"message": "Request sent to the backend"})
	} else {
		valid := utils.IsValidAddress(userReq.ToWallet)
		if valid {
			fmt.Println("Faucet : Send 1 ether to the Specific Address :", userReq.ToWallet)
			faucet.SendTransactionLegacy(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), float64(1))
			c.JSON(200, gin.H{"message": "Request sent to the backend"})
		} else {
			c.JSON(200, gin.H{"message": "Public address format is not valid"})
		}
	}
}

func StartSimulationHandler(c *gin.Context) {
	var simuReq SimuRequest

	if SimuRunning {
		fmt.Println("Simulation already started")
		c.JSON(http.StatusOK, gin.H{"message": "Simulation already started"})
		return
	}

	if err := c.BindJSON(&simuReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if simuReq.AccountsPerWallet < 1 && reflect.TypeOf(simuReq.AccountsPerWallet) != reflect.TypeOf(0) {
		fmt.Println("PERSONNALIZED ERROR : accounts_per_wallets bad parameter")
		return
	} else if simuReq.EthersPerWallet < 1 && reflect.TypeOf(simuReq.EthersPerWallet) != reflect.TypeOf(0) {
		fmt.Println("PERSONNALIZED ERROR : ethers_per_wallets bad parameter")
		return
	} else if simuReq.EthersPerTransaction < 0 && reflect.TypeOf(simuReq.EthersPerTransaction) != reflect.TypeOf(1.0) {
		fmt.Println("PERSONNALIZED ERROR : ethers_per_transactions bad parameter")
		return
	} else if simuReq.TransactionsPerBlock < 2 && reflect.TypeOf(simuReq.TransactionsPerBlock) != reflect.TypeOf(0) {
		fmt.Println("PERSONNALIZED ERROR : transactions_per_blocks bad parameter")
		return
	}
	fmt.Println("Going to start Simulation")
	go Simulation(simuReq, stopChannel)
	SimuRunning = true
	c.JSON(http.StatusOK, gin.H{"message": "Simulation started"})
	fmt.Println("Simulation started")
}

func StopSimulationHandler(c *gin.Context) {
	var simuReq SimuRequest

	if err := c.BindJSON(&simuReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !SimuRunning {
		fmt.Println("Simulation : Simulation already stopped")
		c.JSON(http.StatusOK, gin.H{"message": "Simulation already stopped"})
		return
	}

	fmt.Println("Simulation : Going to stop Simulation")
	SimuRunning = false
	stopChannel <- true
	c.JSON(http.StatusOK, gin.H{"message": "Simulation stopped"})
}

func Simulation(simuReq SimuRequest, stopChan chan bool) {

	richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(config.Connection.Rich_private_key)
	utils.ErrManagement(err)

	clientWs, err = ethclient.Dial(config.Connection.Ws_endpoint)
	utils.ErrManagement(err)

	simulation.Simulation(clientWs, richPrivKey, richPubKey, simuReq.AccountsPerWallet, simuReq.EthersPerWallet, simuReq.EthersPerTransaction, simuReq.TransactionsPerBlock, stopChan)

}
