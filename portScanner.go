package main

import (
	"net"
	"time"
)

type portStatus struct {
	Port   string
	Status string
}

type Tag struct {
	Name  string
	Value string
}

type hostStatus struct {
	Host       string
	Results    []portStatus
	OpenCount  uint
	CloseCount uint
	Duration   uint
	Tags       []Tag
}

func portScan(host string, ports []string, tags []Tag) (status hostStatus) {
	timeout := 100 * time.Millisecond
	status.Host = host
	startTime := uint(time.Now().UnixMilli())

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
		status.Duration = uint(time.Now().UnixMilli()) - startTime
		status.Tags = tags
	}

	return status
}
