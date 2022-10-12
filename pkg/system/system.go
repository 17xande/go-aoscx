package system

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func Get(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "system")
	system, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get system summary: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}

func GetStatus(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "system/status")
	system, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get system status: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}

func GetCPU(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "system/status/cpu")
	system, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get system status cpu: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}

func GetMemory(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "system/status/memory")
	system, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get system status memory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}
