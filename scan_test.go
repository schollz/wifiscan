package wifiscan

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	stdout, _, err := runCommand(1*time.Second, "echo hi")
	assert.True(t, strings.Contains(stdout, "hi"))
	assert.Nil(t, err)
}

func TestScan(t *testing.T) {
	w, err := Scan(os.Getenv("WIFI"))
	assert.Nil(t, err)
	assert.NotEmpty(t, w)
}
