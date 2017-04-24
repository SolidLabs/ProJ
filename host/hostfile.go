package host

import (
	"fmt"
	"github.com/cbednarski/hostess"
	"os"
)

func AddHostEntry(hostNames []string, ip string, port int) {
	hostsfile := loadHostFile()

	for _, host := range hostNames {
		hostname := hostess.NewHostname(host, ip, true)
		replace := hostsfile.Hosts.ContainsDomain(hostname.Domain)

		if replace {
			fmt.Printf("Updating entry %s\n", hostname.FormatHuman())
		} else {
			fmt.Printf("Added %s\n", hostname.FormatHuman())
		}

		hostsfile.Hosts.Add(hostname)
	}

	saveHostFile(hostsfile)
}

func loadHostFile() *hostess.Hostfile {
	hostsfile, errs := hostess.LoadHostfile()

	if len(errs) > 0 {
		for _, err := range errs {
			os.Stderr.WriteString(err.Error())
		}
		os.Stderr.WriteString("Errors while parsing hostsfile")
	}

	return hostsfile
}

func saveHostFile(hostfile *hostess.Hostfile) {
	err := hostfile.Save()
	if err != nil {
		fmt.Println(hostess.ErrCantWriteHostFile.Error())
	}
}
