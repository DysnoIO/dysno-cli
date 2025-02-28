package config

import (
	"github.com/dysnoio/dysno-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show the current configuration",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Println(pterm.Cyan("Current configuration:"))

		// Account address
		var accountSetting string
		if viper.GetString("account") == "" {
			accountSetting = pterm.Gray("No account set")
		} else {
			accountSetting = viper.GetString("account")
		}
		pterm.Println("  - Account address:  ", accountSetting)

		// Contract address
		var caSetting string
		if viper.GetString("ca") == "" {
			caSetting = pterm.Gray("No contract address set")
		} else {
			caSetting = viper.GetString("ca")
		}
		pterm.Println("  - Contract address: ", caSetting)

		// Chain ID
		var chainSetting string
		if viper.GetString("chainId") == "" {
			chainSetting = pterm.Gray("No chain set")
		} else {
			chainSetting = viper.GetString("chainId")
			for name, id := range utils.CHAINS {
				if id.String() == chainSetting {
					chainSetting = name
					break
				}
			}
		}
		pterm.Println("  - Chain:            ", chainSetting)

		// RPC HTTPs
		var rpcHttpSetting string
		if viper.GetString("rpc.http") == "" {
			rpcHttpSetting = pterm.Gray("No RPC HTTPs URL set")
		} else {
			rpcHttpSetting = viper.GetString("rpc.http")
		}
		pterm.Println("  - RPC HTTPs URL:    ", rpcHttpSetting)

		// RPC Websocket
		var rpcWsSetting string
		if viper.GetString("rpc.ws") == "" {
			rpcWsSetting = pterm.Gray("No RPC Websocket URL set")
		} else {
			rpcWsSetting = viper.GetString("rpc.ws")
		}
		pterm.Println("  - RPC Websocket URL:", rpcWsSetting)

		// Gas price
		var gasLimitSetting string
		if viper.GetString("gasLimit") == "" {
			gasLimitSetting = pterm.Gray("No gas limit set")
		} else {
			gasLimitSetting = viper.GetString("gasLimit")
		}
		pterm.Println("  - Gas limit:        ", gasLimitSetting)

	},
}

func init() {}
