package config

import (
	"github.com/spf13/cobra"
)

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Configure the miner RPC and CA settings",
	Long:  "",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	ConfigCmd.AddCommand(rpcCmd)
	ConfigCmd.AddCommand(caCmd)
	ConfigCmd.AddCommand(gaslimitCmd)
	ConfigCmd.AddCommand(chainCmd)
	ConfigCmd.AddCommand(showCmd)
}
