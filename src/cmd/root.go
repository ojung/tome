package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	homedir "github.com/mitchellh/go-homedir"
)

var cfgFile string
var debug bool

var historyFilePathConfigKey = "historyFile"
var shellTypeConfigkey = "shellType"
var repositoryConfigKey = "repository"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "tome",
	Short: "Share shell spells.",
	Long: `Share shell spells with other wizards.
	This application aspires to be a shared zsh/bash history.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.tome{.yaml|.json})")
	rootCmd.PersistentFlags().BoolP("debug", "d", false, "set to true to see stack traces")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".tome" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".tome")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
		requireParam(historyFilePathConfigKey)
		requireParam(shellTypeConfigkey)
		requireParam(repositoryConfigKey)
	}
}

func requireParam(configKey string) {
	if !viper.IsSet(configKey) {
		panic(fmt.Sprintf("Missing required config parameter: %s.", configKey))
	}
}
