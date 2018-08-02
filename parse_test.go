package wifiscan

import (
	"io/ioutil"
	"log"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

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
