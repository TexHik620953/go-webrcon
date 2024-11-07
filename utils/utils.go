package utils

import (
	"net"
)

func IsWsTimeout(err interface{}) bool {
	if netErr, ok := err.(net.Error); ok {
		return netErr.Timeout()
	}
	return false
}
