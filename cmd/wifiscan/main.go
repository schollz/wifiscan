package main

import (
	"flag"
	"fmt"

	log "github.com/schollz/logger"
	wifiscan "github.com/schollz/wifiscan"
)

func main() {
	var wifiInterface string
	var debug bool
	flag.StringVar(&wifiInterface, "wifi", "", "set wifi interface (linux only)")
	flag.BoolVar(&debug, "debug", false, "set debug mode")
	flag.Parse()
	if debug {
		log.SetLevel("debug")
	} else {
		log.SetLevel("info")
	}
	wifis, err := wifiscan.Scan(wifiInterface)
	if err != nil {
		log.Error(err)
	}
	if len(wifis) > 0 {
		fmt.Println("SSID\t\t\tRSSI")
	} else {
		fmt.Println("no mac addresses found")
	}
	for _, w := range wifis {
		fmt.Printf("%s\t%d\n", w.SSID, w.RSSI)
	}
}
