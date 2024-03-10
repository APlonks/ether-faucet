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
	AccountsPerWallets    int     `json:"accounts_per_wallets"`
	EthersPerWallets      int     `json:"ethers_per_wallets"`
	EthersPerTransactions float64 `json:"ethers_per_transactions"`
	TransactionsPerBlocks int     `json:"transactions_per_blocks"`
}

func main() {

	// Initialisation du canal stopChannel
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
		faucet.SendTransactionLegacy(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), float64(config.Simulation.Ethers))
		c.JSON(200, gin.H{"message": "Request sent to the backend"})
	} else {
		fmt.Println("The Address :", userReq.ToWallet)
		valid := utils.IsValidAddress(userReq.ToWallet)
		if valid {
			fmt.Println("Send 1 ether to the Specific Address :", userReq.ToWallet)
			faucet.SendTransactionLegacy(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), float64(config.Simulation.Ethers))
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

	// Lire et valider la requête
	if err := c.BindJSON(&simuReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if simuReq.AccountsPerWallets < 1 && reflect.TypeOf(simuReq.AccountsPerWallets) != reflect.TypeOf(0) {
		fmt.Println("PERSONNALIZED ERROR : accounts_per_wallets bad parameter")
		return
	} else if simuReq.EthersPerWallets < 1 && reflect.TypeOf(simuReq.EthersPerWallets) != reflect.TypeOf(0) {
		fmt.Println("PERSONNALIZED ERROR : ethers_per_wallets bad parameter")
		return
	} else if simuReq.EthersPerTransactions < 0 && reflect.TypeOf(simuReq.EthersPerTransactions) != reflect.TypeOf(1.0) {
		fmt.Println("PERSONNALIZED ERROR : ethers_per_transactions bad parameter")
		return
	} else if simuReq.TransactionsPerBlocks < 2 && reflect.TypeOf(simuReq.TransactionsPerBlocks) != reflect.TypeOf(0) {
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

	fmt.Println("Simulation : Going stop Simulation")
	SimuRunning = false
	stopChannel <- true
	c.JSON(http.StatusOK, gin.H{"message": "Simulation stopped"})
}

func Simulation(simuReq SimuRequest, stopChan chan bool) {

	// Config Client //////////////
	richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(config.Connection.Rich_private_key)
	utils.ErrManagement(err)

	clientWs, err = ethclient.Dial(config.Connection.Ws_endpoint)
	utils.ErrManagement(err)

	simulation.Simulation(clientWs, richPrivKey, richPubKey, simuReq.AccountsPerWallets, simuReq.EthersPerWallets, simuReq.EthersPerTransactions, simuReq.TransactionsPerBlocks, stopChan)

}
