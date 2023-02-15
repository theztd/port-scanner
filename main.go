package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	var inPath, outPath, templateName, customTemplateFile, customPorts, promFile, inTags string
	var pRange []string
	var data []hostStatus

	// default ports have to be scan
	ports := "1-1200,1900,2049,2379-2380,2483,2484,3306,3389,4646,5000-5005,5060,5432,6379,6443,6881,6999,8080,8300,8500,9200,9300,9100,10250,10257,10259,30000-32767"

	flag.StringVar(&inPath, "in", "", "REQUIRED: Path to the file with hosts (One line = one host)")
	flag.StringVar(&outPath, "out", "", "Define the file path where output have to be stored.")
	flag.StringVar(&templateName, "template", "html", "Name of the output template (build in are: json, prometheus, html).")
	flag.StringVar(&customTemplateFile, "template-file", "", "Path to the custom template file.")
	flag.StringVar(&customPorts, "ports", ports, "Custom port definition (Example: \"22,80,443,9100-9200,5432\")")
	flag.StringVar(&promFile, "prom", "", "Save prometheus metric from port scanning to the given file")
	flag.StringVar(&inTags, "tags", "", "Add custom tags (Example: env:prod,filename:source/file.txt)")

	flag.Parse()

	if len(inPath) < 1 {
		fmt.Printf("\nFATAL: Missing arguments!!!\n\nPlease read help and try it again...\n")
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Exit fast, if outPath is defined but not writeble, exit ASAP
	out := os.Stdout
	if len(outPath) > 0 {
		fout, err := os.Create(outPath)
		if err != nil {
			log.Panicln("PANIC: Unable to save output to the given path", err)
		}
		defer fout.Close()
		out = fout
	}

	// Generate list of scanned ports
	if len(customPorts) > 2 {
		pRange = portsToRange(customPorts)
	} else {
		pRange = portsToRange(ports)
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

	// Default tags
	tags := append([]Tag{}, Tag{Name: "filename", Value: filepath.Base(inPath)})
	// Replace default tags if are defined
	if len(inTags) > 3 {
		tags = tagParser(inTags)
	}
	fmt.Println("TAGS:", tags)

	for scan.Scan() {
		// Add wait group counter for each line
		wg.Add(1)

		// Paralel run
		go func(host string) {
			// Finish wait group (countdown wait group)
			defer wg.Done()

			// Run host scanning
			log.Println("INFO: Start scanning ", host)
			data = append(data, portScan(host, pRange, tags))
			log.Println("INFO: Done scanning ", host)
		}(scan.Text())

	}

	// Wait till waitgroup will be 0
	wg.Wait()

	fmt.Println(data)

	// Write prom metrics from measurement if promFile is defined
	if len(promFile) > 1 {
		pOut, err := os.Create(promFile)
		if err != nil {
			log.Println("ERROR: unable to write prometheus metrics, but continue", err)
		}
		defer pOut.Close()
		renderResults(data, "prometheus", "", pOut)
	}

	// Render results
	renderResults(data, templateName, customTemplateFile, out)
}
