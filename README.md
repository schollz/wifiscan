# wifiscan

[![go report card](https://goreportcard.com/badge/github.com/schollz/wifiscan)](https://goreportcard.com/report/github.com/schollz/wifiscan) 
[![coverage](https://cover.run/go/github.com/schollz/wifiscan.svg)](https://gocover.io/github.com/schollz/wifiscan)
[![godocs](https://godoc.org/github.com/schollz/wifiscan?status.svg)](https://godoc.org/github.com/schollz/wifiscan) 


A platform-independent WiFi scanning library for getting BSSID + RSSI from nearby access points. It should work on most Linux installations as well as Windows and OS X.

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
```

You can also use the command-line tool as:

```
$ wifiscan
20:4e:7f:8e:4e:ec: -85
f0:5c:19:a2:26:71: -88
f0:5c:19:a2:28:51: -75
...
```

## Contributing

Pull requests are welcome. Feel free to...

- Revise documentation
- Add new features
- Fix bugs
- Suggest improvements

## License

MIT
