# wifiscan

[![go report card](https://goreportcard.com/badge/github.com/schollz/wifiscan)](https://goreportcard.com/report/github.com/schollz/wifiscan) 
[![coverage](https://cover.run/go/github.com/schollz/wifiscan.svg)](https://gocover.io/github.com/schollz/wifiscan)
[![godocs](https://godoc.org/github.com/schollz/wifiscan?status.svg)](https://godoc.org/github.com/schollz/wifiscan) 


A platform-independent WiFi scanning library for getting BSSID + RSSI from nearby access points. It should work on most Linux installations as well as Windows and OS X.


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
```

## Contributing

Pull requests are welcome. Feel free to...

- Revise documentation
- Add new features
- Fix bugs
- Suggest improvements

## License

MIT
