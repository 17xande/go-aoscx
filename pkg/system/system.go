package system

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func Get(ip string) {
	r := request.New(ip, "system")
	system, err := r.Get()
	if err != nil {
		fmt.Printf("can't get system summary: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}

func GetStatus(ip string) {
	r := request.New(ip, "system/status")
	system, err := r.Get()
	if err != nil {
		fmt.Printf("can't get system status: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}

func GetCPU(ip string) {
	r := request.New(ip, "system/status/cpu")
	system, err := r.Get()
	if err != nil {
		fmt.Printf("can't get system status cpu: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}

func GetMemory(ip string) {
	r := request.New(ip, "system/status/memory")
	system, err := r.Get()
	if err != nil {
		fmt.Printf("can't get system status memory: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(system)
}
