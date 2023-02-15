package main

import (
	"strconv"
	"strings"
)

func tagParser(tagsIn string) (tags []Tag) {
	for _, tPair := range strings.Split(tagsIn, ",") {
		k_v := strings.Split(tPair, ":")
		tags = append(tags, Tag{Name: k_v[0], Value: k_v[1]})
	}
	return tags
}

func portsToRange(customPorts string) (portList []string) {
	/*
		Convert string like "80,443,1000-1005,9100"

		to

		[]strings{80,443,1000,1001,1002,1003,1004,1005,9100}

	*/
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
