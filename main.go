package main

import (
	"flag"
	"fmt"
	"strings"
)

var (
	configFile string
	config     Config
	mode       string
	debug      = true
)

func init() {
	flag.StringVar(&configFile, "c", "config.json", "Config file")
	flag.StringVar(&mode, "m", "", "Mode")
	flag.BoolVar(&debug, "d", false, "Debug flag")
	flag.Parse()
}

func main() {
	if configFile == "" {
		fmt.Println("No config specified.")
		return
	} else {
		loadConfig()
	}
	LogMsg("Config: %+v\n", config)
	if strings.Contains("discord", mode) {
		startDiscord()
	} else if strings.Contains("api", mode) {
		startRest()
	} else {
		fmt.Printf("%+v is not a recognized mode.\n", mode)
	}
}

func LogMsg(format string, a ...interface{}) {
	if !debug {
		return
	}
	fmt.Printf(fmt.Sprintf("%+v\n", format), a...)
}
