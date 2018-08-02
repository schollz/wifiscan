package main

import (
	"fmt"
	"log"

	wifiscan "github.com/schollz/wifi-scan"
)

func main() {
	wifis, err := wifiscan.Scan()
	if err != nil {
		log.Fatal(err)
	}
	for _, w := range wifis {
		fmt.Printf("%s: %d\n", w.SSID, w.RSSI)
	}
}
