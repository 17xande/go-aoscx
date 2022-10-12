package vlan

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/internal/request"
)

func GetAll(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "vlans")
	vlans, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get vlans: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(vlans)
}
