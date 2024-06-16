package wifiscan

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Wifi is the data structure containing the basic
// elements
type Wifi struct {
	SSID  string `json:"ssid"`
	BSSID string `json:"bssid"`
	RSSI  int    `json:"rssi"`
}

// Parse will parse wifi output and extract the access point
// information.
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
		if w.BSSID == "" {
			if strings.Contains(line, "BSSID") {
				fs := strings.Fields(line)
				if len(fs) == 4 {
					w.BSSID = fs[3]
				}
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
		if strings.Contains(line, "SSID") && !strings.Contains(line, "BSSID") {
			w.SSID = strings.Trim(line[strings.Index(line, ":")+1:], " ")
		}
		if w.BSSID != "" && w.RSSI != 0 {
			wifis = append(wifis, w)
			w.BSSID = ""
			w.RSSI = 0
		}
	}

	return
}

func parseDarwin(output string) (wifis []Wifi, err error) {
	const _MAC_ADDRESS_REGEX string = "([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})"

	scanner := bufio.NewScanner(strings.NewReader(output))
	for scanner.Scan() {
		line := strings.Trim(scanner.Text(), " ")

		// To check if it's a valid line, check if there's a MAC address in it
		var valid_line bool
		if valid_line, err = regexp.MatchString(_MAC_ADDRESS_REGEX, line); err != nil || !valid_line {
			continue
		}

		var mac_regex *regexp.Regexp = regexp.MustCompile(_MAC_ADDRESS_REGEX)
		var mac_loc []int = mac_regex.FindStringIndex(line)

		var ssid string = strings.Trim(line[:mac_loc[0]], " ")
		var bssid string = strings.Trim(line[mac_loc[0]:mac_loc[1]], " ")
		var rssi int
		rssi, err = strconv.Atoi(strings.Trim(line[mac_loc[1]:mac_loc[1]+4], " "))

		wifis = append(wifis, Wifi{SSID: ssid, BSSID: bssid, RSSI: rssi})
	}

	return
}

func parseLinux(output string) (wifis []Wifi, err error) {
	scanner := bufio.NewScanner(strings.NewReader(output))
	w := Wifi{}
	for scanner.Scan() {
		line := scanner.Text()
		if w.BSSID == "" {
			if strings.Contains(line, "Address") {
				fs := strings.Fields(line)
				if len(fs) == 5 {
					w.BSSID = strings.ToLower(fs[4])
				}
			}
		} else {
			if strings.Contains(line, "Signal level=") {
				var level int
				level, err = strconv.Atoi(strings.Split(strings.Split(strings.Split(line, "level=")[1], "/")[0], " dB")[0])
				if err != nil {
					continue
				}
				if level > 0 {
					level = (level / 2) - 100
				}
				w.RSSI = level
			}
		}
		if strings.Contains(line, "ESSID") {
			var trimmed_line string = strings.Trim(line, " ")
			w.SSID = strings.Trim(trimmed_line[strings.Index(trimmed_line, ":")+1:], " \"")
		}
		if w.BSSID != "" && w.RSSI != 0 {
			wifis = append(wifis, w)
			w.BSSID = ""
			w.RSSI = 0
		}
	}

	return
}
