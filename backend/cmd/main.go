package main

import (
	"crypto/ecdsa"
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

func main() {

	var (
		clientHttp  *ethclient.Client
		clientWs    *ethclient.Client
		err         error
		choice      int
		richPrivKey *ecdsa.PrivateKey
		richPubKey  common.Address
		config      *utils.Config
	)

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

	clientWs, err = ethclient.Dial(config.Connection.Ws_endpoint)
	utils.ErrManagement(err)

	log.Println("Starting the HTTP server on port 9090")
	router := mux.NewRouter()

	router.HandleFunc("/sendTransactions", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			// Gérer l'erreur
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}
		// Supposons que le champ du formulaire pour l'adresse du portefeuille est nommé "walletAddress"
		walletAddress := r.FormValue("walletAddress")

		// Ici, vous appelleriez votre fonction avec `walletAddress` comme argument
		// Assurez-vous que `faucet.SendEthersToSpecificWallet` accepte les bons paramètres
		faucet.SendEthersToSpecificAddress(clientHttp, richPrivKey, richPubKey, walletAddress, 1)

		// Envoyer une réponse de succès
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Ethers sent successfully")
	}).Methods("POST") // Assurez-vous de spécifier la méthode HTTP si nécessaire

	http.ListenAndServe(":9090", nil)
}
