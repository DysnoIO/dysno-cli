package mine

import (
	"context"
	"math/big"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/dysnoio/dysno-cli/token"
	"github.com/dysnoio/dysno-cli/utils"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var key *keystore.Key

var MineCmd = &cobra.Command{
	Use:   "mine",
	Short: "Mine nuggets to collect Dysno tokens",
	Long:  "",
	PreRun: func(cmd *cobra.Command, args []string) {
		// Check if default account is set
		if viper.GetString("account") == "" {
			pterm.Error.Println("No default account set.")
			pterm.Println()
			pterm.Println(pterm.Cyan("If you don't have an account, use the ") + pterm.Magenta("`account add`") + pterm.Cyan(" command to create one."))
			pterm.Info.Println("Adding an account will also set it as default.")
			pterm.Println()
			pterm.Println(pterm.Cyan("Use the ") + pterm.Magenta("`account list`") + pterm.Cyan(" command to list available accounts."))
			pterm.Println(pterm.Cyan("Use the ") + pterm.Magenta("`account set`") + pterm.Cyan(" command to set an available account as default."))
			os.Exit(0)
		}

		// Find default account
		ks := keystore.NewKeyStore(viper.GetString("keystore"), keystore.StandardScryptN, keystore.StandardScryptP)
		accountKs, err := ks.Find(accounts.Account{Address: common.HexToAddress(viper.GetString("account"))})
		if err != nil {
			pterm.Error.Println("Error finding default account keystore file.")
			os.Exit(0)
		}

		// Read default account keystore file
		jsonBytes, err := os.ReadFile(accountKs.URL.Path)
		if err != nil {
			pterm.Error.Println("Error reading the default account keystore file.")
			os.Exit(0)
		}

		// Load default account
		pterm.Info.Println("Default account:", pterm.Green(accountKs.Address.Hex()))
		pterm.Println("Enter keystore encryption password." + pterm.Gray(" (*** It will be masked. ***)"))
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Password")

		key, err = keystore.DecryptKey(jsonBytes, password)
		if err != nil {
			pterm.Println()
			pterm.Error.Println("Error loading default account. Likely wrong password.")
			os.Exit(0)
		}

		// Show success message
		pterm.Println()
		pterm.Success.Println("Loaded default account successfully!")
		pterm.Println()

	},
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to the network using the http and ws RPC endpoints
		clientHttp, err := ethclient.Dial(viper.GetString("rpc.http"))
		if err != nil {
			pterm.Error.Println("Failed to connect to the network.")
			os.Exit(0)
		}
		clientWs, err := ethclient.Dial(viper.GetString("rpc.ws"))
		if err != nil {
			pterm.Error.Println("Failed to connect to the network.")
			os.Exit(0)
		}

		// Connect to the contract
		address := common.HexToAddress(viper.GetString("ca"))
		instance, err := token.NewToken(address, clientHttp)
		if err != nil {
			pterm.Error.Println("Failed to create contract instance.")
			os.Exit(0)
		}

		// Parse contract ABI
		contractAbi, err := abi.JSON(strings.NewReader(token.TokenABI))
		if err != nil {
			pterm.Error.Println("Failed to parse contract ABI.")
			os.Exit(0)
		}

		// Define event topics to filter logs
		mineEventTopic := crypto.Keccak256Hash([]byte("Mine(address,uint256,uint256,bytes32)"))
		readjustEventTopic := crypto.Keccak256Hash([]byte("Readjust(uint256,uint256)"))

		// Define filter query
		query := ethereum.FilterQuery{
			Addresses: []common.Address{address},
			Topics:    [][]common.Hash{{mineEventTopic, readjustEventTopic}},
		}

		// Define mining variables outside the loop
		sender := common.HexToAddress(key.Address.Hex())
		nonce := new(big.Int).SetUint64(0)
		hash := common.Hash{}
		var wg sync.WaitGroup

		// Create channels
		stopChan := make(chan struct{})
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		logsChan := make(chan types.Log)

		// Subscribe to logs
		sub, err := clientWs.SubscribeFilterLogs(context.Background(), query, logsChan)
		if err != nil {
			pterm.Error.Println("Failed to subscribe to logs.")
			os.Exit(0)
		}

		// Get mining parameters
		challengeNumber, target, err := utils.MiningParams(instance)
		if err != nil {
			pterm.Error.Println("Failed to get mining parameters.")
			os.Exit(0)
		}

		// Start spinner
		pterm.DefaultSpinner.WithShowTimer(false).Start("Mining... " + pterm.Gray("Press Ctrl/Cmd + C to stop."))

		// Goroutine to run the loop
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-sub.Err():
					// Sunbscription error
					pterm.Error.Println("Failed to subscribe to logs. Ensure the RPC endpoints and contract address are correct.")
					os.Exit(0)
				case vLog := <-logsChan:
					// Handle logs
					switch vLog.Topics[0] {
					case mineEventTopic: // Mine event
						event := struct {
							To           common.Address
							Value        *big.Int
							Epoch        *big.Int
							NewChallenge [32]byte
						}{}
						err := contractAbi.UnpackIntoInterface(&event, "Mine", vLog.Data)
						if err != nil {
							pterm.Error.Println("Failed to unpack log data.")
							os.Exit(0)
						}
						challengeNumber = event.NewChallenge
						nonce.SetUint64(0)
						pterm.Println("Event: Challenge number updated.")
					case readjustEventTopic: // Readjust target event
						event := struct {
							Timestamp *big.Int
							NewTarget *big.Int
						}{}
						err := contractAbi.UnpackIntoInterface(&event, "Readjust", vLog.Data)
						if err != nil {
							pterm.Error.Println("Failed to unpack log data.")
							os.Exit(0)
						}
						target = event.NewTarget
						pterm.Println("Event: Mining target updated.")
					}
				case <-stopChan:
					// Stop signal received
					pterm.Info.Println("Mining stopping...")
					return
				default:
					// CORE MINING LOGIC
					// Calculate hash
					hash = crypto.Keccak256Hash(
						common.BytesToHash(challengeNumber[:]).Bytes(),
						sender.Bytes(),
						nonce.FillBytes(make([]byte, 32)),
					)
					hashAsBigInt := new(big.Int).SetBytes(hash.Bytes())

					if hashAsBigInt.Cmp(target) <= 0 {
						// Hash is a nugget candidate
						tx, err := utils.MineNugget(key, clientHttp, instance, nonce)
						if err != nil {
							// Mining transaction failed. Refetch mining parameters
							pterm.Warning.Println("Mine transaction failed. Refetching mining parameters...\nError:", err)
							challengeNumber, target, err = utils.MiningParams(instance)
							if err != nil {
								pterm.Error.Println("Failed to get mining parameters.")
								os.Exit(0)
							}
							pterm.Info.Println("Mining parameters refetched.")
						} else {
							// Mining transaction successful
							pterm.Println()
							pterm.Success.Println("Nugget Mined!\nTransaction ID:", pterm.Green(tx.Hash().Hex()))
							pterm.Println()
						}
						// Reset nonce to 0 after attempting to mine a nugget
						nonce.SetUint64(0)
					} else {
						// Not a nugget candidate. Increment nonce
						nonce.Add(nonce, big.NewInt(1))
					}
				}
			}
		}()

		// Wait for a signal
		<-signalChan
		pterm.Println()
		pterm.Warning.Println("Received stop signal.")
		close(stopChan)

		// Wait for the mining routine to finish
		wg.Wait()
		pterm.Success.Println("Mining stopped gracefully.")
	},
}

func init() {}
