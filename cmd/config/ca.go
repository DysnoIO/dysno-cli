package config

import (
	"github.com/dysnoio/dysno-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var caCmd = &cobra.Command{
	Use:   "ca",
	Short: "Set the contract address",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Prompt for the contract address and save it to the config file
		ca, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(utils.DEFAULT_CA).Show("Token contract address")
		viper.Set("ca", ca)
		viper.WriteConfig()

		pterm.Println()
		pterm.Success.Println("Updated the contract address.")

		pterm.Println()
		pterm.Warning.Println("Be sure to test your connection settings with the", pterm.Magenta("`test`"), "command.")
	},
}

func init() {}
