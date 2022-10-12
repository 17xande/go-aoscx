package ports

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func GetAll(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "ports")
	ports, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get all ports: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ports)
}

func Get(ip string, ID int) {
	url := fmt.Sprintf(request.UrlAPI, ip, "ports/"+fmt.Sprint(ID))
	ports, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get all ports: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ports)
}
