package wifiscan

import (
	"bufio"
	"fmt"
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
		wifis, err = parseDarwin(output)
	case "linux":
		wifis, err = parseLinux(output)
	default:
		err = fmt.Errorf("%s is not a recognized OS", os)
	}
	return
}

func parseWindows(output string) (wifis []Wifi, err error) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	w := Wifi{}
	wifis = []Wifi{}
	for scanner.Scan() {
		line := scanner.Text()
		if w.SSID == "" {
			if strings.Contains(line, "BSSID") {
				fs := strings.Fields(line)
				if len(fs) == 4 {
					w.SSID = fs[3]
				}
			} else {
				continue
			}
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

func parseDarwin(output string) (wifis []Wifi, err error) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	wifis = []Wifi{}
	for scanner.Scan() {
		line := scanner.Text()
		fs := strings.Fields(line)
		if len(fs) < 6 {
			continue
		}
		rssi, errParse := strconv.Atoi(fs[2])
		if errParse != nil {
			continue
		}
		if rssi > 0 {
			continue
		}
		wifis = append(wifis, Wifi{SSID: strings.ToLower(fs[1]), RSSI: rssi})
	}
	return
}

func parseLinux(output string) (wifis []Wifi, err error) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	w := Wifi{}
	wifis = []Wifi{}
	for scanner.Scan() {
		line := scanner.Text()
		if w.SSID == "" {
			if strings.Contains(line, "Address") {
				fs := strings.Fields(line)
				if len(fs) == 5 {
					w.SSID = strings.ToLower(fs[4])
				}
			} else {
				continue
			}
		} else {
			if strings.Contains(line, "Signal level=") {
				level, errParse := strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "level=")[1], "/")[0], " dB")[0])
				if errParse != nil {
					continue
				}
				if level > 0 {
					level = (level / 2) - 100
				}
				w.RSSI = level
			}
		}
		if w.SSID != "" && w.RSSI != 0 {
			wifis = append(wifis, w)
			w = Wifi{}
		}
	}
	return
}
