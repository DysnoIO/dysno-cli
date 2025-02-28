package account

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set an account as default",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Load account options
		ks := keystore.NewKeyStore(viper.GetString("keystore"), keystore.StandardScryptN, keystore.StandardScryptP)
		var options []string
		for _, account := range ks.Accounts() {
			options = append(options, account.Address.Hex())
		}

		// Show warning if no accounts are found
		if len(options) == 0 {
			pterm.Warning.Println("No accounts found.")
			os.Exit(0)
		}

		// Show current default account
		if viper.GetString("account") != "" {
			pterm.Info.Println("Current default account: " + pterm.Green(viper.GetString("account")))
		}

		pterm.Println()
		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Select an account to set as default")

		// Set the selected account as default
		viper.Set("account", selectedOption)
		err := viper.WriteConfig()
		if err != nil {
			pterm.Error.Println("Error writing to config file.")
			os.Exit(0)
		}

		// Display the selected option
		pterm.Println()
		pterm.Success.Printfln("Successfully set default account - %s", pterm.Green(selectedOption))
	},
}

func init() {}
