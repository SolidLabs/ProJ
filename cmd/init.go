package cmd

import (
	"github.com/solidworx/proj/host"
	"github.com/solidworx/proj/webserver"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

var hostNames []string
var ip string
var port int

func init() {
	RootCmd.AddCommand(initCmd)

	initCmd.Flags().StringArrayVarP(&hostNames, "host", "n", nil, "Project hostnames")
	initCmd.Flags().IntVarP(&port, "port", "p", 80, "Hostname port")
	initCmd.Flags().StringVarP(&ip, "ip", "i", "127.0.0.1", "IP address to use for hostname")
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a local config",
	Long:  `Initialize a local config`,
	Run: func(cmd *cobra.Command, args []string) {
		webserver.AddConfig(hostNames, ip, port, getProjectName())
		host.AddHostEntry(hostNames, ip, port)
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
