package config

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rpcCmd = &cobra.Command{
	Use:   "rpc",
	Short: "Set the RPC URLs",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for the RPC URLs and save them to the config file
		pterm.Println("Obtain RPC URLs from a blockchain node provider.")
		pterm.Println("(e.g. Alchemy, Infura, Quicknode)")

		pterm.Println()
		http, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(viper.GetString("rpc.http")).Show("HTTPs URL")
		viper.Set("rpc.http", http)
		ws, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(viper.GetString("rpc.ws")).Show("Websocket URL")
		viper.Set("rpc.ws", ws)
		viper.WriteConfig()

		pterm.Println()
		pterm.Success.Println("RPC URLs set successfully!")
		pterm.Warning.Println("Be sure to test your connection settings with the", pterm.Magenta("`test`"), "command.")
	},
}

func init() {}
