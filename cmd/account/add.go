package account

import (
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an account",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Println("Enter account private key." + pterm.Gray(" (*** It will be masked. ***)"))
		result, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Private key")

		// Result must be a valid private key
		privateKey, err := crypto.HexToECDSA(result)
		if err != nil {
			pterm.Println()
			pterm.Error.Println("Invalid private key.")
			os.Exit(0)
		}

		// Import account
		pterm.Println()
		pterm.Println("Enter keystore encryption password." + pterm.Gray(" (*** It will be masked. ***)"))
		password, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Password")
		retypePassword, _ := pterm.DefaultInteractiveTextInput.WithMask("*").Show("Retype password")

		if password != retypePassword {
			pterm.Println()
			pterm.Error.Println("Passwords do not match.")
			os.Exit(0)
		}

		ks := keystore.NewKeyStore(viper.GetString("keystore"), keystore.StandardScryptN, keystore.StandardScryptP)
		account, err := ks.ImportECDSA(privateKey, password)
		if err != nil {
			pterm.Println()
			pterm.Error.Println("Failed to import account. It may already exist.")
			os.Exit(0)
		}

		pterm.Println()
		pterm.Success.Println("Account added successfully -", account.Address.Hex())

		// Set account as default if no default account is set
		if viper.GetString("account") == "" {
			viper.Set("account", account.Address.Hex())
			viper.WriteConfig()
			pterm.Println()
			pterm.Info.Println("Account set as default.")
		}
	},
}

func init() {}
