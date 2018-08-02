package wifiscan

import (
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunCommand(t *testing.T) {
	stdout, stderr, err := runCommand(1*time.Second, "ls -lSh")
	log.Println(stdout, stderr)
	assert.Nil(t, err)
}
