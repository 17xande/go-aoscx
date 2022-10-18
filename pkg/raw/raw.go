package raw

import (
	"fmt"
	"os"

	"github.com/17xande/go-aoscx/pkg/request"
)

func Get(ip, path string) {
	r := request.New(ip, path)
	raw, err := r.Get()
	if err != nil {
		fmt.Printf("can't get system summary: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(raw)
}
