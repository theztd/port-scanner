package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	log.Println("INFO: Application started")

	var pRange []string
	var data []hostStatus
	var inPath, outPath string
	var quick, full bool

	flag.StringVar(&inPath, "in", "", "Path to the file with hosts (One line = one host)")
	flag.StringVar(&outPath, "out", "", "Path to the output file.")
	flag.BoolVar(&quick, "quick", false, "Do only fast scan (predefined most common ports)")
	flag.BoolVar(&full, "full", false, "Scan everything from 1-65535 (Super slow)")
	flag.Parse()

	if quick {
		for _, p := range []int{21, 22, 23, 25, 53, 80, 110, 123, 143, 389, 443} {
			pRange = append(pRange, strconv.Itoa(p))
		}
	} else if full {
		for i := 1; i < 65536; i++ {
			pRange = append(pRange, strconv.Itoa(i))
		}
	} else {
		for i := 21; i < 1024; i++ {
			pRange = append(pRange, strconv.Itoa(i))
		}
	}

	if !full {
		for _, p := range []int{1900, 2049, 3306, 3389, 4646, 5000, 5001, 5004, 5005, 5060, 6379, 8080, 8300, 8500, 9999, 5432, 9200, 9300, 9100} {
			pRange = append(pRange, strconv.Itoa(p))
		}
	}

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
	renderResults(data)
}
