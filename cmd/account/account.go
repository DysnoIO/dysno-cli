package account

import (
	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Manage accounts",
	Long:  `Manage accounts. You can add, delete, and list accounts. Use the set command to set an account as default.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	AccountCmd.AddCommand(addCmd)
	AccountCmd.AddCommand(deleteCmd)
	AccountCmd.AddCommand(listCmd)
	AccountCmd.AddCommand(setCmd)
}
