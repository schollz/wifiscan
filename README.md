# wifiscan

[![go report card](https://goreportcard.com/badge/github.com/schollz/wifiscan)](https://goreportcard.com/report/github.com/schollz/wifiscan) 
[![coverage](https://img.shields.io/badge/coverage-87%25-brightgreen.svg)](https://gocover.io/github.com/schollz/wifiscan)
[![godocs](https://godoc.org/github.com/schollz/wifiscan?status.svg)](https://godoc.org/github.com/schollz/wifiscan) 


A platform-independent WiFi scanning library for getting SSID + BSSID + RSSI from nearby access points. It should work on most Linux installations as well as Windows and OS X.

## How does it work?

*wifiscan* works by calling the OS-specific Wifi scan utility and parsing the output. For Linux this is `iwlist`, for Windows it is `netsh.exe` and for OS X it is `airport`. Other systems are not supported at the moment (although, as long as they are Linux-based I believe you can install `iwlist`).

## Install

```
go get -u github.com/schollz/wifiscan/...
```

## Usage 

You can use it in your Go code as:

```golang
wifis, err := wifiscan.Scan()
if err != nil {
    log.Fatal(err)
}
for _, w := range wifis {
    fmt.Println(w.SSID, w.BSSID, w.RSSI)
}
```

You can also use the command-line tool as:

```
$ WIFI=wlan0 wifiscan
SSID            BSSID                   RSSI
NOS-618F        2c:a1:7d:69:61:8f       -98
MEO-WiFi        1c:57:3e:28:07:62       -97
MEO-WiFi        1c:57:3e:2b:de:12       -93
MEO-WiFi        58:fc:20:b1:af:a2       -88
...
```

_Note:_ When using with Linux you will need to add `sudo` to get a full scan.

## Contributing

Pull requests are welcome. Feel free to...

- Revise documentation
- Add new features
- Fix bugs
- Suggest improvements

## License

MIT
