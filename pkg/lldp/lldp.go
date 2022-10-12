package lldp

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func GetGlobalStatus(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "lldp")
	lldp, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get LLDP Global Status: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(lldp)
}

func GetLocalPorts(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "lldp/local-port")
	lldp, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get LLDP local ports config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(lldp)
}

func GetLocalPort(ip string, ID int) {
	url := fmt.Sprintf(request.UrlAPI, ip, "lldp/local-port/"+fmt.Sprint(ID))
	lldp, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get LLDP local port config: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(lldp)
}
