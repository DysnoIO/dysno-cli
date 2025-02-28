package config

import (
	"github.com/dysnoio/dysno-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var chainCmd = &cobra.Command{
	Use:   "chain",
	Short: "Set the chain ID",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Info.Println("Current chain ID:", viper.GetString("chainId"))

		options := []string{}
		for name := range utils.CHAINS {
			options = append(options, name)
		}

		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Select a chain")

		viper.Set("chainId", utils.CHAINS[selectedOption])
		viper.WriteConfig()

		pterm.Println()
		pterm.Success.Println("Updated the chain to:", selectedOption)
	},
}

func init() {}
