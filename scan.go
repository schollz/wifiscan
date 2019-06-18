package wifiscan

import (
	"bytes"
	"fmt"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	log "github.com/schollz/logger"
)

var TimeLimit = 10 * time.Second

func init() {
	log.SetLevel("info")
}

// Scan will scan the optional interface for wifi access points
func Scan(wifiInterface ...string) (wifilist []Wifi, err error) {
	if runtime.GOOS == "linux" && (len(wifiInterface) == 0 || wifiInterface[0] == "") {
		var interfaces []string
		interfaces, err = getInterfacesLinux()
		if err != nil {
			log.Debug(err)
			return
		}
		for _, in := range interfaces {
			var w []Wifi
			w, err = scan(in)
			if len(w) > 0 {
				wifilist = append(wifilist, w...)
			}
		}
		if len(wifilist) > 0 {
			err = nil
			wifimap := make(map[string]Wifi)
			for _, w := range wifilist {
				wifimap[w.SSID] = w
			}
			i := 0
			for _, w := range wifimap {
				wifilist[i] = w
				i++
			}
			wifilist = wifilist[:i]
		}

		return
	}
	return scan()
}

func scan(wifiInterface ...string) (wifilist []Wifi, err error) {
	command := ""
	os := ""
	switch runtime.GOOS {
	case "windows":
		os = "windows"
		command = "netsh.exe wlan show networks mode=Bssid"
		_, _, errRun := runCommand(TimeLimit, "netsh interface set interface name=Wi-Fi admin=disabled")
		if errRun != nil {
			log.Debug(errRun)
		}
		_, _, errRun = runCommand(TimeLimit, "netsh interface set interface name=Wi-Fi admin=enabled")
		if errRun != nil {
			log.Debug(errRun)
		}
		time.Sleep(3 * time.Second)
	case "darwin":
		os = "darwin"
		command = "/System/Library/PrivateFrameworks/Apple80211.framework/Versions/Current/Resources/airport -s"
	default:
		os = "linux"
		command = "iwlist wlan0 scan"
		if len(wifiInterface) > 0 && len(wifiInterface[0]) > 0 {
			command = fmt.Sprintf("iwlist %s scan", wifiInterface[0])
		}
	}
	stdout, _, err := runCommand(TimeLimit, command)
	if err != nil {
		log.Debugf("error on '%s': %s", command, err)
		log.Debug(stdout)
		return
	}
	wifilist, err = Parse(stdout, os)
	return
}

func runCommand(tDuration time.Duration, commands string) (stdout, stderr string, err error) {
	log.Debugf("running '%s' for %s", commands, tDuration)
	command := strings.Fields(commands)
	cmd := exec.Command(command[0])
	if len(command) > 0 {
		cmd = exec.Command(command[0], command[1:]...)
	}
	var outb, errb bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &errb
	err = cmd.Start()
	if err != nil {
		return
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case <-time.After(tDuration):
		err = cmd.Process.Kill()
	case err = <-done:
		stdout = outb.String()
		stderr = errb.String()
	}
	return
}

func getInterfacesLinux() (interfaces []string, err error) {
	stdout, _, err := runCommand(TimeLimit, "ip address")
	if err != nil {
		return
	}
	return getInteracesFromString(stdout)
}

func getInteracesFromString(s string) (interfaces []string, err error) {
	for _, line := range strings.Split(s, "\n") {
		if !strings.Contains(line, "BROADCAST") {
			continue
		}
		cols := strings.Split(line, ":")
		if len(cols) < 3 {
			continue
		}
		_, errConvert := strconv.Atoi(cols[0])
		if errConvert != nil {
			continue
		}
		if strings.Contains(cols[1], "@") || strings.Contains(cols[1], "docker") {
			continue
		}
		interfaces = append(interfaces, strings.TrimSpace(cols[1]))
	}
	return
}
