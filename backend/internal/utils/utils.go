package utils

import (
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

// IsValidAddress validate hex address
func IsValidAddress(iaddress interface{}) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	switch v := iaddress.(type) {
	case string:
		return re.MatchString(v)
	case common.Address:
		return re.MatchString(v.Hex())
	default:
		return false
	}
}

// Not used
func ErrManagement(err error) {
	if err != nil {
		// log.Fatal("!! ERROR !!:", err)
		fmt.Println("!! ERROR !!:", err)
	}
}
