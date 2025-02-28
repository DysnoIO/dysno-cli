package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/dysnoio/dysno-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var gaslimitCmd = &cobra.Command{
	Use:   "gaslimit",
	Short: "Set the gas limit for CLI transactions",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {

		// Show the recommended gas limit
		pterm.Info.Println("Recommended gas limit:", utils.DEFAULT_GAS_LIMIT)

		// Prompt for the gas limit and save it to the config file
		gasLimit, _ := pterm.DefaultInteractiveTextInput.WithDefaultValue(fmt.Sprintf("%d", viper.GetUint64("gasLimit"))).Show("Gas limit")

		// Parrse the gas limit to uint64 and save it to the config file
		parsedGasLimit, err := strconv.ParseUint(gasLimit, 10, 64)
		if err != nil {
			pterm.Error.Println("Invalid gas limit. Please enter a valid number.")
			os.Exit(0)
		}

		viper.Set("gasLimit", parsedGasLimit)
		viper.WriteConfig()

		pterm.Println()
		pterm.Success.Println("Updated the gas limit for CLI transactions.")
	},
}

func init() {}
