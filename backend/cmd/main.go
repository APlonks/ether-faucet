package main

import (
	"backend/internal/faucet"
	"backend/internal/utils"
	"backend/internal/wallets"
	"crypto/ecdsa"
	"fmt"
	"net/http"

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
	choice      int
	richPrivKey *ecdsa.PrivateKey
	richPubKey  common.Address
)

func main() {

	configPath, err = utils.ParseFlags()
	config, err = utils.LoadConfig(configPath)

	router := gin.Default()

	router.Use(cors.Default()) // Allow all

	router.GET("/someGet", getting)

	router.POST("/SendEthersToSpecificAddress", SendEthersToSpecificAddress)
	router.Run() // listen and serve on 0.0.0.0:8080
}

func getting(c *gin.Context) {
	fmt.Println(c)
	fmt.Println("TEST")
}

type UserRequest struct {
	ToWallet string `json:"wallet"`
}

func SendEthersToSpecificAddress(c *gin.Context) {

	// Config Client //////////////
	richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(config.Connection.Rich_private_key)
	utils.ErrManagement(err)

	clientHttp, err = ethclient.Dial(config.Connection.Http_endpoint)
	utils.ErrManagement(err)
	//////////////////////////////

	var userReq UserRequest

	if err := c.BindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userReq.ToWallet == "" {
		fmt.Println("Send 1 ether to the Specific Address : 0x0000000000000000000000000000000000000000")
		faucet.SendEthersToSpecificAddress(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), config.Simulation.Ethers)
	} else {
		fmt.Println("Send 1 ether to the Specific Address :", userReq.ToWallet)
		faucet.SendEthersToSpecificAddress(clientHttp, richPrivKey, richPubKey, common.HexToAddress(userReq.ToWallet), config.Simulation.Ethers)
	}

	c.JSON(200, gin.H{"status": "Request sent to the backend"})

}
