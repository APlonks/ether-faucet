package main

import (
	"backend/internal/faucet"
	"backend/internal/simulation"
	"backend/internal/utils"
	"backend/internal/wallets"
	"crypto/ecdsa"
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	RICH_PRIVATE_KEY string
	HTTP_ENDPOINT    string
	WS_ENDPOINT      string
	err              error
	clientHttp       *ethclient.Client
	clientWs         *ethclient.Client
	richPrivKey      *ecdsa.PrivateKey
	richPubKey       common.Address
	stopChannel      chan bool // Simulation control
	SimuRunning      bool
)

type SimuRequest struct {
	AccountsPerWallet    int     `json:"accounts_per_wallet"`
	EthersPerWallet      int     `json:"ethers_per_wallet"`
	EthersPerTransaction float64 `json:"ethers_per_transaction"`
	TransactionsPerBlock int     `json:"transactions_per_block"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Print("Error:", err)
	}

	RICH_PRIVATE_KEY = os.Getenv("RICH_PRIVATE_KEY")
	HTTP_ENDPOINT = os.Getenv("HTTP_ENDPOINT")
	WS_ENDPOINT = os.Getenv("WS_ENDPOINT")

	// Init channel stopChannel
	stopChannel = make(chan bool)

	router := gin.Default()

	router.Use(cors.Default()) // Allow all

	router.GET("/testing", Testing)

	router.POST("/faucet", SendEthersToSpecificAddress)

	router.POST("/start-simulation", StartSimulationHandler)

	router.POST("/stop-simulation", StopSimulationHandler)

	router.Run(":5002") // Port 8080 by default
}

func Testing(c *gin.Context) {
	fmt.Println()
	c.JSON(http.StatusOK, gin.H{"message": "API connected"})
}

func SendEthersToSpecificAddress(c *gin.Context) {
	fmt.Println()
	type UserRequest struct {
		ToWallet string `json:"wallet"`
	}
	var userReq UserRequest

	richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(RICH_PRIVATE_KEY)
	if err != nil {
		fmt.Println("Error while trying to retrieve Public key from Private key:", err)
		return
	}

	clientHttp, err = ethclient.Dial(HTTP_ENDPOINT)
	if err := c.BindJSON(&userReq); err != nil {
		fmt.Println("Error while trying to dial:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("The client:", clientHttp)
	fmt.Println("The privateKey:", richPrivKey)
	fmt.Println("The fromAddres:", richPubKey)
	fmt.Println("The toWallet:", common.HexToAddress(userReq.ToWallet))
	if userReq.ToWallet == "" {
		fmt.Println("Faucet : Send 1 ether to the Specific Address : 0x0000000000000000000000000000000000000000")
		err = faucet.SendTransactionLegacy(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), float64(1))
		if err != nil {
			fmt.Println("Error while trying to send Transaction Legacy")
			c.JSON(200, gin.H{"message": "Cannot send transaction to the blockchain"})
			return
		}

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
	fmt.Println()
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

	// wg.Add(1)
	go func() {
		fmt.Println("Going to start Simulation")
		err = Simulation(simuReq, stopChannel)
		if err != nil {
			c.JSON(200, gin.H{"message": "Failed to start simulation due to internal error.", "details": err.Error()})
			fmt.Println("Going out from anonymous")
		}
	}()

	SimuRunning = true
	c.JSON(http.StatusOK, gin.H{"message": "Simulation started"})
	fmt.Println("Simulation started")
}

func StopSimulationHandler(c *gin.Context) {
	fmt.Println()
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

func Simulation(simuReq SimuRequest, stopChan chan bool) error {
	fmt.Println()

	richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(RICH_PRIVATE_KEY)
	if err != nil {
		fmt.Println("Error while retrieving public key from private key:", err)
		return err
	}

	clientWs, err = ethclient.Dial(WS_ENDPOINT)
	if err != nil {
		fmt.Println("Error during setup with the WS_ENDPOINT:", err)
		return err
	}

	err = simulation.Simulation(clientWs, richPrivKey, richPubKey, simuReq.AccountsPerWallet, simuReq.EthersPerWallet, simuReq.EthersPerTransaction, simuReq.TransactionsPerBlock, stopChan)
	if err != nil {
		fmt.Println("Error while trying to execute Simulation:", err)
		return err
	}

	return nil
}
