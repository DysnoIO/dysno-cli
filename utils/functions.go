package utils

import (
	"context"
	"errors"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/viper"

	"github.com/dysnoio/dysno-cli/token"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pterm/pterm"
)

func MiningParams(instance *token.Token) ([32]byte, *big.Int, error) {
	challengeNumber, err := instance.GetChallengeNumber(nil)
	if err != nil {
		return [32]byte{}, nil, errors.New("error getting challenge number")
	}
	target, err := instance.GetMiningTarget(nil)
	if err != nil {
		return [32]byte{}, nil, errors.New("error getting mining target")
	}
	return challengeNumber, target, nil
}

func MineNugget(key *keystore.Key, client *ethclient.Client, instance *token.Token, nonce *big.Int) (*types.Transaction, error) {
	txNonce, err := client.PendingNonceAt(context.Background(), key.Address)
	if err != nil {
		return nil, errors.New("error getting transaction nonce")
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, errors.New("error getting suggested gas price")
	}
	chainID, ok := new(big.Int).SetString(viper.GetString("chainId"), 10)
	if !ok {
		return nil, errors.New("invalid chain ID")
	}
	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		return nil, errors.New("error creating transaction signer")
	}

	auth.Nonce = big.NewInt(int64(txNonce))
	auth.Value = big.NewInt(0)                  // in wei
	auth.GasLimit = viper.GetUint64("gasLimit") // in units
	auth.GasPrice = gasPrice

	tx, err := instance.Mine(auth, nonce)
	if err != nil {
		return nil, errors.New("nugget mining conditions not met")
	}
	return tx, nil
}

func WorkingDir() string {
	wd, err := os.Getwd()
	if err != nil {
		pterm.Error.Println("Error getting the working directory.")
		os.Exit(0)
	}
	return wd
}
