package main

import (
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"

	"backend/internal/faucet"
	"backend/internal/utils"
	"backend/internal/wallets"
)

// Middleware pour configurer les en-têtes CORS
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Ajustez ceci pour des origines spécifiques si nécessaire
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		// Si c'est une requête preflight, nous envoyons seulement les en-têtes CORS avec un statut 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Appel au prochain middleware ou routeur
		next.ServeHTTP(w, r)
	})
}

func main() {

	var (
		clientHttp *ethclient.Client
		// clientWs    *ethclient.Client
		err         error
		richPrivKey *ecdsa.PrivateKey
		richPubKey  common.Address
		config      *utils.Config
	)

	type WalletRequest struct {
		Wallet string `json:"wallet"`
	}

	// Generate our config based on the config supplied
	// by the user in the flags
	configPath, err := utils.ParseFlags()
	utils.ErrManagement(err)

	config, err = utils.LoadConfig(configPath)
	utils.ErrManagement(err)

	richPrivKey, richPubKey, err = wallets.RetrieveKeysFromHexHashedPrivateKey(config.Connection.Rich_private_key)
	utils.ErrManagement(err)

	clientHttp, err = ethclient.Dial(config.Connection.Http_endpoint)
	utils.ErrManagement(err)

	// clientWs, err = ethclient.Dial(config.Connection.Ws_endpoint)
	// utils.ErrManagement(err)

	log.Println("Starting the HTTP server on port 9999")
	router := mux.NewRouter()

	// Applies the CORS middleware
	router.Use(corsMiddleware)

	router.HandleFunc("/SendEthersToSpecificAddress", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("w:", w)
		fmt.Println("r:", r)

		// Creates an instance of WalletRequest to store the decoded data
		var request WalletRequest

		// Decodes the JSON request body into the request instance
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		// Now, it can use request.Wallet as the wallet address
		walletAddress := request.Wallet
		fmt.Println("The walletAddress:", walletAddress)

		// Here, it would call your function with `walletAddress` as the argument
		// Ensures that `faucet.SendEthersToSpecificWallet` accepts the correct parameters
		faucet.SendEthersToSpecificAddress(clientHttp, richPrivKey, richPubKey, common.HexToAddress(walletAddress), 1000)

		// Sends a success response
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Ethers sent successfully")
	}).Methods("POST", "OPTIONS") // Ensures to specify the HTTP method if necessary

	http.ListenAndServe(":9999", router)
}
