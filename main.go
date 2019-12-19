package main

import (
	"fmt"
	"os"
	"runtime"
)

var config map[string]interface{}
var servicesPath string

//var startedService []os.Process

func main() {
	fmt.Printf("%sp%si%sn%si%st%s (c) 2018-2019 Maksim Pinigin\n", COLOR_LIGHT_GREEN, COLOR_LIGHT_RED, COLOR_LIGHT_CYAN, COLOR_LIGHT_PURPLE, COLOR_YELLOW, COLOR_RESET)

	var loadConfResult bool
	switch runtime.GOOS {
	case "linux", "freebsd":
		loadConfResult = LoadConfig("/etc/pinit/configuration.json")
	case "windows":
		loadConfResult = LoadConfig("C:\\pinit\\configuration.json")
	default:
		FatalError("Unsupported operating system. Built on " + runtime.GOOS, nil)
		os.Exit(3)
	}

	if loadConfResult != true {
		fmt.Println("pinit: Fatal error: unknown error loadConfResult = false")
		os.Exit(2)
	}

	servicesPath = config["services_path"].(string)

	Init()

	select { }
}
