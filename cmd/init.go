package cmd

import (
	"github.com/solidworx/proj/host"
	"github.com/solidworx/proj/webserver"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type HostConfig struct {
	HostNames []string
	Ip        string
	Port      int
}

var hostConfig = &HostConfig{}

func init() {

	RootCmd.AddCommand(initCmd)

	initCmd.Flags().StringArrayVarP(&hostConfig.HostNames, "host", "n", nil, "Project hostnames")
	initCmd.Flags().IntVarP(&hostConfig.Port, "port", "p", 80, "Hostname port")
	initCmd.Flags().StringVarP(&hostConfig.Ip, "ip", "i", "127.0.0.1", "IP address to use for hostname")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a local config",
	Long:  `Initialize a local config`,
	Run: func(cmd *cobra.Command, args []string) {
		webserver.AddConfig(hostConfig, getProjectName())
		host.AddHostEntry(hostConfig)
	},
}

func getProjectName() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	dirs := strings.Split(dir, string(os.PathSeparator))

	return dirs[len(dirs)-1]
}
