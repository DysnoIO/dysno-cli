package utils

import (
	"math/big"
	"path/filepath"
)

var DEFAULT_CA = "0xe18F76E8193F81cE2F18e9536e6E32461Fd30A5A"

var DEFAULT_GAS_LIMIT = uint64(200000)
var DEFAULT_CHAIN_ID = big.NewInt(8453)

var DEFAULT_CONFIG_FILE_NAME = ".dysno-cli"
var DEFAULT_CONFIG_FILE_EXT = "yaml"
var DEFAULT_KEYSTORE_DIR = "keystore"

var DEFAULT_CONFIG_PATH = filepath.Join(WorkingDir(), DEFAULT_CONFIG_FILE_NAME+"."+DEFAULT_CONFIG_FILE_EXT)
var DEFAULT_KEYSTORE_PATH = filepath.Join(WorkingDir(), DEFAULT_KEYSTORE_DIR)

var CHAINS = map[string]*big.Int{
	"Base (Chain ID: 8453)":          big.NewInt(8453),
	"Base Sepolia (Chain ID: 84532)": big.NewInt(84532),
}
