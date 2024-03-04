package utils

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Connection struct {
		Rich_private_key string `yaml:"rich_private_key"`
		Http_endpoint    string `yaml:"http_endpoint"`
		Ws_endpoint      string `yaml:"ws_endpoint"`
	} `yaml:"connection"`
	Simulation struct {
		Accounts     int `yaml:"accounts"`
		Ethers       int `yaml:"ethers"`
		Transactions int `yaml:"transactions"`
	} `yaml:"simulation"`
}

func LoadConfig(configPath string) (*Config, error) {
	var c Config
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatalf("Error reading YAML file: %s\n", err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Error parsing YAML file: %s\n", err)
	}
	fmt.Println("OK parsing YAML file")

	return &c, nil
}

func ParseFlags() (string, error) {
	// String that contains the configured configuration path
	var configPath string

	// Set up a CLI flag called "-config" to allow users
	// to supply the configuration file
	flag.StringVar(&configPath, "config", "./config.yml", "path to config file")

	// Actually parse the flags
	flag.Parse()

	// Validate the path first
	// if err := ValidateConfigPath(configPath); err != nil {
	// 	return "", err
	// }

	// Return the configuration path
	return configPath, nil
}

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

func ErrManagement(err error) {
	if err != nil {
		log.Fatal("!! ERROR !!:", err)
	}
}
