package wallets

import (
	"crypto/ecdsa"
	"fmt"
	"hash"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

type Wallet struct {
	Key        ecdsa.PrivateKey
	KeyHex     string
	Address    common.Address
	AddressHex string
}

// Create a wallet by calloing CreateWallet() function
func NewWallets() Wallet {
	return CreateWallet()
}

func CreateWallet() Wallet {

	var (
		privateKey     *ecdsa.PrivateKey
		err            error
		publicKeyECDSA *ecdsa.PublicKey
		ok             bool
		publicKeyBytes []byte
		address        common.Address
		hash           hash.Hash
	)

	fmt.Println()
	privateKey, err = crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("Private key in *ecdsa.PrivateKey format", privateKey)

	privateKeyBytes := crypto.FromECDSA(privateKey)
	// fmt.Println("Private key in hex format:", hexutil.Encode(privateKeyBytes)[:]) // 0xfad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19

	publicKey := privateKey.Public()
	publicKeyECDSA, ok = publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	publicKeyBytes = crypto.FromECDSAPub(publicKeyECDSA)
	// fmt.Println("EDCSA public key in Slice of bytes",hexutil.Encode(publicKeyBytes)[4:]) // 9a7df67f79246283fdc93af76d4f8cdd62c4886e8cd870944e817dd0b97934fdd7719d0810951e03418205868a5c1b40b192451367f28e0088dd75e15de40c05

	address = crypto.PubkeyToAddress(*publicKeyECDSA)
	// fmt.Println("Public key in string:", address)                             // 0x96216849c49358B10257cb55b28eA603c874b05E
	// fmt.Println("Public key in hexadecimal (common.Address):", address.Hex()) // 0x96216849c49358B10257cb55b28eA603c874b05E

	hash = sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes[1:])
	// fmt.Println("Public key in hexadecimal hashed:", hexutil.Encode(hash.Sum(nil)[12:]))

	wallet := Wallet{Key: *privateKey, KeyHex: hexutil.Encode(privateKeyBytes), Address: address, AddressHex: hexutil.Encode(hash.Sum(nil)[12:])}
	return wallet
}

func CreateWallets(numWallets int) []Wallet {

	wallets := make([]Wallet, numWallets)

	for i := 0; i < numWallets; i++ {
		wallets[i] = NewWallets()
	}

	return wallets
}

func RetrieveKeysFromHexHashedPrivateKey(privateKeyHex string) (*ecdsa.PrivateKey, common.Address, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	return privateKey, fromAddress, nil
}
