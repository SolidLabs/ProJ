package config

import "github.com/spf13/afero"

type HostConfig struct {
	HostNames []string
	Ip        string
	Port      int
}

var Fs = afero.NewOsFs()
