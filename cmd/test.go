package cmd

import (
	"os"

	"github.com/dysnoio/dysno-cli/token"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Test connection using the RPC and CA settings",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to the network
		client, err := ethclient.Dial(viper.GetString("rpc.http"))
		if err != nil {
			pterm.Error.Println("Failed to connect to the network. Check the RPC settings.")
			os.Exit(0)
		}

		// Connect to the contract
		address := common.HexToAddress(viper.GetString("ca"))
		instance, err := token.NewToken(address, client)
		if err != nil {
			pterm.Error.Println("Failed to create contract instance.")
			os.Exit(0)
		}

		// Get the token name
		name, err := instance.Name(nil)
		if err != nil {
			pterm.Info.Println("Update the RPC and CA settings in the config file.")
			os.Exit(0)
		}

		// Display the token name as a success message
		pterm.Success.Println("Connected to the", name, "contract.")
	},
}

func init() {}
