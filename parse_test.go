package wifiscan

import (
	"log"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkParseWindows(b *testing.B) {
	pathToTest := "testing/netsh"
	files, err := os.ReadDir(pathToTest)
	if err != nil {
		log.Fatal(err)
	}

	var output string
	for _, f := range files {
		bs, _ := os.ReadFile(path.Join("testing/netsh", f.Name()))
		output = string(bs)
		break
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Parse(output, "windows")
	}
}

func TestParseWindows(t *testing.T) {
	pathToTest := "testing/netsh"
	files, err := os.ReadDir(pathToTest)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		b, err := os.ReadFile(path.Join(pathToTest, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		ws, err := Parse(string(b), "windows")
		assert.Nil(t, err)
		assert.NotEmpty(t, ws)
	}
}

func TestParseOSX(t *testing.T) {
	pathToTest := "testing/airport"
	files, err := os.ReadDir(pathToTest)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		b, err := os.ReadFile(path.Join(pathToTest, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		ws, err := Parse(string(b), "darwin")
		assert.Nil(t, err)
		assert.NotEmpty(t, ws)
	}
}

func TestParseLinux(t *testing.T) {
	pathToTest := "testing/iwlist"
	files, err := os.ReadDir(pathToTest)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		b, err := os.ReadFile(path.Join(pathToTest, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		ws, err := Parse(string(b), "linux")
		assert.Nil(t, err)
		assert.NotEmpty(t, ws)
	}
}
