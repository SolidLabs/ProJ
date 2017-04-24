package cmd

import (
	"github.com/solidworx/proj/host"
	"github.com/solidworx/proj/webserver"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"github.com/solidworx/proj/config"
	"sync"
)

var hostConfig = &config.HostConfig{}

const fnLength = 2

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
		flist := [fnLength]func(){}

		flist[0] = func() {
			webserver.AddConfig(hostConfig, getProjectName())
		}

		flist[1] = func() {
			host.AddHostEntry(hostConfig)
		}

		var wg sync.WaitGroup
		wg.Add(len(flist))
		go pool(&wg, len(flist), flist)

		wg.Wait()
	},
}

func worker(tasksCh <-chan func(), wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasksCh
		if !ok {
			return
		}

		task()
	}
}

func pool(wg *sync.WaitGroup, workers int, tasks [fnLength]func()) {
	tasksCh := make(chan func())

	for i := 0; i < workers; i++ {
		go worker(tasksCh, wg)
	}

	for i := 0; i < len(tasks); i++ {
		tasksCh <- tasks[i]
	}

	close(tasksCh)
}

func getProjectName() string {
	dir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	dirs := strings.Split(dir, string(os.PathSeparator))

	return dirs[len(dirs)-1]
}
