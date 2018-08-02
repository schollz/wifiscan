package wifiscan

import "fmt"

type WiFi struct {
	SSID string `json:"ssid"`
	RSSI int    `json:"rssi"`
}

func Parse(output, os string) (wifis []WiFi, err error) {
	switch os {
	case "windows":

	case "darwin":

	case "linux":

	default:
		err = fmt.Errorf("%s is not a recognized OS", os)
	}
	return
}

func parseWindows(output string) (wifis []WiFi, err error) {
	return
}
