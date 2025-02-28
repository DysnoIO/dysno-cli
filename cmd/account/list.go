package account

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all accounts",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		// Load account options
		ks := keystore.NewKeyStore(viper.GetString("keystore"), keystore.StandardScryptN, keystore.StandardScryptP)
		var options []string
		for _, account := range ks.Accounts() {
			options = append(options, account.Address.Hex())
		}

		if len(options) == 0 {
			pterm.Println()
			pterm.Warning.Println("No accounts found.")
			os.Exit(0)
		}

		// Show accounts
		defaultFound := false
		pterm.Println(pterm.Cyan("Accounts:"))
		for _, account := range ks.Accounts() {
			if viper.GetString("account") == account.Address.Hex() {
				// Show default account with a different color
				pterm.Println(pterm.Green("  "+account.Address.Hex()), pterm.Gray(" (default)"))
				defaultFound = true
				continue
			}
			pterm.Println(pterm.Cyan("  " + account.Address.Hex()))
		}

		// Show warning if no default account is set
		if !defaultFound {
			pterm.Println()
			pterm.Warning.Println("Default account not found. Use the " + pterm.Magenta("`account set`") + " command to set an account as default.")
		}
	},
}

func init() {}
