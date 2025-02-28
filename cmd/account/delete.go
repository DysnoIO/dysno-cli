package account

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an account",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		ks := keystore.NewKeyStore(viper.GetString("keystore"), keystore.StandardScryptN, keystore.StandardScryptP)

		var options []string
		for _, account := range ks.Accounts() {
			options = append(options, account.Address.Hex())
		}

		if len(options) == 0 {
			pterm.Warning.Println("No accounts found.")
			os.Exit(0)
		}

		selectedOption, _ := pterm.DefaultInteractiveSelect.WithOptions(options).Show("Select an account to delete")

		// Show selected option
		pterm.Println()
		pterm.Info.Printfln("Selected option: %s", pterm.Green(selectedOption))

		// Show confirmation
		pterm.Println()
		result, _ := pterm.DefaultInteractiveConfirm.Show("Are you sure you want to delete this account?")

		if result {
			// Find the account to delete
			accountToDelete, err := ks.Find(accounts.Account{Address: common.HexToAddress(selectedOption)})
			if err != nil {
				pterm.Println()
				pterm.Error.Println("Error finding account to delete.")
				os.Exit(0)
			}

			// Ask for the encryption password
			pterm.Println()
			pterm.Println("Enter keystore encryption password." + pterm.Gray(" (*** It will be masked. ***)"))
			password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Password")

			err = ks.Delete(accountToDelete, password)
			if err != nil {
				pterm.Println()
				pterm.Error.Println("Error deleting account. Probably wrong password.")
				os.Exit(0)
			}

			// Update the config file if the deleted account was the default account
			if viper.GetString("account") == accountToDelete.Address.Hex() {
				viper.Set("account", "")
				viper.WriteConfig()
				pterm.Println()
				pterm.Warning.Println("Default account deleted. Use the " + pterm.Magenta("`account set`") + " command to set a new default account.")
			}

			// Print a success message
			pterm.Println()
			pterm.Success.Println("Account deleted successfully!")
		} else {
			// Print an error message
			pterm.Println()
			pterm.Error.Println("Account deletion cancelled.")
		}
	},
}

func init() {}
