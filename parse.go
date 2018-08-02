package wifiscan

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Wifi struct {
	SSID string `json:"ssid"`
	RSSI int    `json:"rssi"`
}

func Parse(output, os string) (wifis []Wifi, err error) {
	switch os {
	case "windows":
		wifis, err = parseWindows(output)
	case "darwin":

	case "linux":

	default:
		err = fmt.Errorf("%s is not a recognized OS", os)
	}
	return
}

var macRegex, _ = regexp.Compile(`(?:[A-Fa-f0-9]{2}[:-]){5}(?:[A-Fa-f0-9]{2})`)

func parseWindows(output string) (wifis []Wifi, err error) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	w := Wifi{}
	wifis = []Wifi{}
	for scanner.Scan() {
		line := scanner.Text()
		if w.SSID == "" {
			mac := macRegex.FindString(line)
			if mac != "" {
				w.SSID = mac
			}
			continue
		} else {
			if strings.Contains(line, "%") {
				fs := strings.Fields(line)
				if len(fs) == 3 {
					w.RSSI, err = strconv.Atoi(strings.Replace(fs[2], "%", "", 1))
					if err != nil {
						return
					}
					w.RSSI = (w.RSSI / 2) - 100
				}
			}
		}
		if w.SSID != "" && w.RSSI != 0 {
			wifis = append(wifis, w)
			w = Wifi{}
		}
	}
	return
}
