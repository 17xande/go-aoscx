package vlan

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func GetAll(ip string) {
	r := request.New(ip, "vlans")
	vlans, err := r.Get()
	if err != nil {
		fmt.Printf("can't get vlans: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(vlans)
}
