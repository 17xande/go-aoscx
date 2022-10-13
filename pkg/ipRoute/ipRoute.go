package ipRoute

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func GetAll(ip string) {
	r := request.New(ip, "ip-route")
	ipRoutes, err := r.Get()
	if err != nil {
		fmt.Printf("can't get all ipRoutes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ipRoutes)
}
