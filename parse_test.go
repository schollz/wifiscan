package wifiscan

import (
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkParseWindows(b *testing.B) {
	pathToTest := "testing/netsh"
	files, err := ioutil.ReadDir(pathToTest)
	if err != nil {
		log.Fatal(err)
	}

	var output string
	for _, f := range files {
		bs, _ := ioutil.ReadFile(path.Join("testing/netsh", f.Name()))
		output = string(bs)
		break
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Parse(output, "windows")
	}
}

func TestParseWindows(t *testing.T) {
	pathToTest := "testing/netsh"
	files, err := ioutil.ReadDir(pathToTest)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		b, err := ioutil.ReadFile(path.Join(pathToTest, f.Name()))
		if err != nil {
			log.Fatal(err)
		}
		ws, err := Parse(string(b), "windows")
		assert.Nil(t, err)
		assert.NotEmpty(t, ws)
	}
}
