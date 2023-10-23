package main

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

func Check[T string](destination T, port T) string {

	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", string(address), timeout)
	var status string
	if err != nil {

		status = color.RedString(fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v", destination, err))

	} else {

		status = color.GreenString(fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v", destination,
			conn.LocalAddr(),
			conn.RemoteAddr()))
	}
	return status
}
