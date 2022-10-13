package ports

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func GetAll(ip string) {
	r := request.New(ip, "ports")
	ports, err := r.Get()
	if err != nil {
		fmt.Printf("can't get all ports: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ports)
}

func Get(ip string, ID int) {
	r := request.New(ip, "ports/"+fmt.Sprint(ID))
	ports, err := r.Get()
	if err != nil {
		fmt.Printf("can't get all ports: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ports)
}
