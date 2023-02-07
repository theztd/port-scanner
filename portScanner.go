package main

import (
	"net"
	"time"
)

type portStatus struct {
	Port   string
	Status string
}

type hostStatus struct {
	Host       string
	Results    []portStatus
	OpenCount  uint
	CloseCount uint
}

func portScan(host string, ports []string) (status hostStatus) {
	timeout := 100 * time.Millisecond
	status.Host = host

	for _, port := range ports {

		// try connection
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)

		s := portStatus{Port: port}
		if err != nil {
			status.CloseCount++
			s.Status = err.Error()
			//log.Println("ERR:", err)
		}

		if conn != nil {
			defer conn.Close()
			s.Status = "Open"
			status.OpenCount++
		}

		status.Results = append(status.Results, s)
	}

	return status
}
