package cmd

import (
	"expvar"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	configFile         string
	isLocalEnvironment bool
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:     "user-management",
	Short:   "",
	Version: "",
}

// SetRevision inject version from git
func SetRevision(r string) {
	if len(r) > 0 {
		RootCmd.Version = fmt.Sprintf("%v-%v", RootCmd.Version, r)
	}
	expvar.NewString("service_version").Set(RootCmd.Version)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&configFile, "config", "config.toml", "config file (default is config.toml)")

	RootCmd.PersistentFlags().BoolVar(&isLocalEnvironment, "local", false, "flag enable local environment")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if configFile != "" { // enable ability to specify config file via flag
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config") // name of config file (without extension)
	}

	viper.AddConfigPath(".") // adding home directory as first search path
	viper.AutomaticEnv()     // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
