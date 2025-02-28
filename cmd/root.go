package cmd

import (
	"os"

	"github.com/dysnoio/dysno-cli/cmd/account"
	"github.com/dysnoio/dysno-cli/cmd/config"
	"github.com/dysnoio/dysno-cli/cmd/mine"
	"github.com/dysnoio/dysno-cli/utils"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "dysno-cli",
	Short: "A brief description of your application",
	Long: pterm.Cyan(`      ____                           
     / __ \ __  __ _____ ____   ____ 
    / / / // / / // ___// __ \ / __ \
   / /_/ // /_/ /(__  )/ / / // /_/ /
  /_____/ \__, //____//_/ /_/ \____/ 
         /____/                      `),
	Version: "1.0.0",
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(0)
	}
}

func init() {
	rootCmd.AddCommand(account.AccountCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(mine.MineCmd)
	rootCmd.AddCommand(testCmd)

	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	// Search config in wd directory with name DEFAULT_CONFIG_FILE_NAME (without extension)
	viper.AddConfigPath(utils.WorkingDir())
	viper.SetConfigName(utils.DEFAULT_CONFIG_FILE_NAME)

	// Read in environment variables that match
	viper.AutomaticEnv()

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err != nil {
		pterm.Warning.Println("No config file found. Creating one...")

		// Create a new config file.
		viper.Set("account", "")
		viper.Set("keystore", utils.DEFAULT_KEYSTORE_PATH)
		viper.Set("rpc.http", "")
		viper.Set("rpc.ws", "")
		viper.Set("ca", utils.DEFAULT_CA)
		viper.Set("gasLimit", utils.DEFAULT_GAS_LIMIT)
		viper.Set("chainId", utils.DEFAULT_CHAIN_ID)
		viper.WriteConfigAs(utils.DEFAULT_CONFIG_PATH)
		pterm.Success.Println("Config file created at:", utils.DEFAULT_CONFIG_PATH)
	}

	// Check that keystore directory is set
	keystoreDir := viper.GetString("keystore")
	if keystoreDir == "" {
		viper.Set("keystore", utils.DEFAULT_KEYSTORE_PATH)
		viper.WriteConfig()
		keystoreDir = utils.DEFAULT_KEYSTORE_PATH
	}

	// Check if keystore directory exists
	if _, err := os.Stat(keystoreDir); os.IsNotExist(err) {
		// Create keystore directory
		err := os.Mkdir(keystoreDir, 0755)
		if err != nil {
			pterm.Error.Println("Error creating keystore directory.")
			os.Exit(0)
		}
	}

	// Check that ca is set
	ca := viper.GetString("ca")
	if ca == "" {
		viper.Set("ca", utils.DEFAULT_CA)
		viper.WriteConfig()
	}

	// Check that gasLimit is set
	gasLimit := viper.GetUint64("gasLimit")
	if gasLimit == 0 {
		viper.Set("gasLimit", utils.DEFAULT_GAS_LIMIT)
		viper.WriteConfig()
	}

	// Check that chainId is set
	chainId := viper.Get("chainId")
	if chainId == nil {
		viper.Set("chainId", utils.DEFAULT_CHAIN_ID)
		viper.WriteConfig()
	}
}
