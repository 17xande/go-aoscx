package main

import (
	"flag"

	"github.com/17xande/go-aoscx/internal/vlan"
)

func main() {
	ip := flag.String("ip", "", "the switche's IP address.")
	command := flag.String("command", "", "a command.")
	flag.Parse()

	if *ip == "" {
		println("no IP address provided.")
		return
	}

	switch *command {
	case "":
		println("no command provided")
	case "vlan":
		vlan.GetAll(*ip)
	default:
		println("command not supported")
	}
}
