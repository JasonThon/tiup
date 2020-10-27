package utils

import (
	"github.com/tj/assert"
	"testing"
)

const ExternalIP = "192.168.0.101"

func TestGetIpAddr(t *testing.T) {
	ip, err := GetExternalIpAddr()
	if err != nil {
		panic("Network broken down")
	}
	assert.Equal(t, ExternalIP , ip.String())
}
