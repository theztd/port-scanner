package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
)

var inPath, outPath, templateName, customTemplateFile, customPorts string

//var quick, full bool

func main() {
	var pRange []string
	var data []hostStatus

	// default ports have to be scan
	ports := "1-1200,1900,2049,2379-2380,2483,2484,3306,3389,4646,5000-5005,5060,5432,6379,6443,6881,6999,8080,8300,8500,9200,9300,9100,10250,10257,10259,30000-32767"

	flag.StringVar(&inPath, "in", "", "Path to the file with hosts (One line = one host)")
	flag.StringVar(&outPath, "out", "", "Path to the output file.")
	flag.StringVar(&templateName, "template", "html", "Name of the output template (build in are: json, prometheus, html).")
	flag.StringVar(&customTemplateFile, "template-file", "", "Path to the custom template file.")
	flag.StringVar(&customPorts, "ports", ports, "Custom port definition (example \"22,80,443,9100-9200,5432\")")

	flag.Parse()

	//
	if len(customPorts) > 2 {
		pRange = customPortsToRange(customPorts)
	} else {
		pRange = customPortsToRange(ports)
	}

	if len(inPath) < 1 {
		fmt.Printf("\nFATAL: Missing arguments!!!\n\nPlease read help and try it again...\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	log.Println("INFO: Application started")
	log.Println("INFO: Scanning ports:", pRange)

	inFile, err := os.Open(inPath)
	if err != nil {
		log.Panicln("ERR:", err)
	}
	defer inFile.Close()

	// Start reading file
	scan := bufio.NewScanner(inFile)

	// Initialize WaitGroup
	var wg sync.WaitGroup

	for scan.Scan() {
		// Add wait group counter for each line
		wg.Add(1)

		// Paralel run
		go func(host string) {
			// Finish wait group (countdown wait group)
			defer wg.Done()

			// Run host scanning
			log.Println("INFO: Start scanning ", host)
			data = append(data, portScan(host, pRange))
			log.Println("INFO: Done scanning ", host)
		}(scan.Text())

	}

	// Wait till waitgroup will be 0
	wg.Wait()

	// Render results
	renderResults(data, customTemplateFile)
}
