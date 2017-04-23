package cmd

import (
    "github.com/spf13/cobra"
    "github.com/spf13/viper"
    "fmt"
    "os"
)

var Vebose bool = false
var cfgFile string

var RootCmd = &cobra.Command{
    Short: "Proj simplifies local development config",
    Long:  ``,
}

func init() {
    cobra.OnInitialize(initConfig)
    RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "$HOME/.proj.yaml", "Path to config file")

    RootCmd.PersistentFlags().BoolVarP(&Vebose, "verbose", "v", false, "Enable verbose logs");
}

func initConfig() {
    if cfgFile != "" {
        viper.SetConfigFile(cfgFile)
    }

    viper.SetConfigName(".proj") // name of config file (without extension)
    viper.AddConfigPath("$HOME") // adding home directory as first search path
    viper.AutomaticEnv()         // read in environment variables that match
    viper.SetEnvPrefix("proj")

    // If a config file is found, read it in.
    if err := viper.ReadInConfig(); err == nil {
        fmt.Println("Using config file:", viper.ConfigFileUsed())
    }
}

func Execute() {
    if err := RootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(-1)
    }
}
