package main

import (
	"strconv"
	"strings"
)

func customPortsToRange(customPorts string) (portList []string) {
	_ports := strings.Split(customPorts, ",")
	for _, p := range _ports {
		if strings.Contains(p, "-") {
			pRange := strings.Split(p, "-")
			start, _ := strconv.Atoi(pRange[0])
			end, _ := strconv.Atoi(pRange[1])
			for i := start; i <= end; i++ {
				portList = append(portList, strconv.Itoa(i))
			}

		} else {
			portList = append(portList, p)
		}
	}

	return portList
}
