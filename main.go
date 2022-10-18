package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/17xande/go-aoscx/pkg/ipRoute"
	"github.com/17xande/go-aoscx/pkg/lldp"
	"github.com/17xande/go-aoscx/pkg/ports"
	"github.com/17xande/go-aoscx/pkg/raw"
	"github.com/17xande/go-aoscx/pkg/system"
	"github.com/17xande/go-aoscx/pkg/vlan"
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
	case "ipRoute":
		ipRoute.GetAll(*ip)
	case "lldp":
		lldp.GetGlobalStatus(*ip)
	case "lldpPorts":
		if len(flag.Args()) == 0 {
			lldp.GetLocalPorts(*ip)
			break
		}

		portID := getTailInt(flag.Args())
		lldp.GetLocalPort(*ip, portID)
	case "ports":
		if len(flag.Args()) == 0 {
			ports.GetAll(*ip)
			break
		}

		portID := getTailInt(flag.Args())
		ports.Get(*ip, portID)
	case "raw":
		if len(flag.Args()) == 0 {
			fmt.Println("no path provided")
			os.Exit(1)
		}

		path := flag.Args()[0]
		raw.Get(*ip, path)
	case "system":
		system.Get(*ip)
	case "systemStatus":
		system.GetStatus(*ip)
	case "systemCPU":
		system.GetCPU(*ip)
	case "systemMemory":
		system.GetMemory(*ip)
	case "vlans":
		vlan.GetAll(*ip)
	default:
		println("command not supported")
	}
}

func getTailInt(args []string) int {
	num, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("last argument must be an integer")
		os.Exit(1)
	}

	return num
}
