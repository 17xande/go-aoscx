package ipRoute

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func GetAll(ip string) {
	url := fmt.Sprintf(request.UrlAPI, ip, "ip-route")
	ipRoutes, err := request.Get(url)
	if err != nil {
		fmt.Printf("can't get all ipRoutes: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(ipRoutes)
}
