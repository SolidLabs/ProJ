package webserver

import (
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"github.com/spf13/afero"
	"fmt"
	"github.com/solidworx/proj/templates"
	"strings"
	cfg "github.com/solidworx/proj/config"
)

func AddConfig(config *cfg.HostConfig, projectDir string) {
	var config_path string = viper.GetString("webservers.nginx.config_path")

	fmt.Println("Writing config to" + config_path)

	var fs afero.File
	fs, _ = cfg.Fs.Create(fmt.Sprintf("%s/%s.conf", cast.ToString(config_path), projectDir))

	_, err := fs.WriteString(fmt.Sprintf(templates.PhpFpmDefault, strings.Join(addPortToHost(config.HostNames, config.Port), " "), projectDir))

	if err != nil {
		fmt.Println(err.Error())
	}
}

func addPortToHost(hosts []string, port int) []string {
	y := hosts[:0]
	for _, n := range hosts {
		y = append(y, n+":"+cast.ToString(port))
	}

	return y
}
