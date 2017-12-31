package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/ivopetiz/portscanner"
)

// PrintBanner presents a simple banner just to identify program name.
func PrintBanner() {

	fmt.Println("\n +----------------+")
	fmt.Println(" | Go PortScanner |")
	fmt.Println(" +----------------+")
}

// main corresponds to main func of portscanner app.
func main() {

	var ports_list []string
	
	start := time.Now()

	ports := flag.String("p", "80", "Port or ports to scan")
	all := flag.Bool("A", false, "Scans from port 1 to 1024")
	flag.Parse()

	if *all {
		*ports = "1-1024"
	}

	ports_list = append(ports_list, *ports)

	PrintBanner()

	// if there's not valid IPs to scan, system will exit with error.

	_ = portscanner.IPScanner(flag.Args(), ports_list, true)

	elapsed := time.Since(start)
	fmt.Println("\nScanned in", elapsed)
}
