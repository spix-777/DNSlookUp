package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// Loggers
var (
	trueLogger  *log.Logger
	falseLogger *log.Logger
)

// Initialize loggers to stdout
func init() {
	trueLogger = log.New(os.Stdout, " [ + ] ", 0)
	falseLogger = log.New(os.Stdout, " [ ! ] ", 0)

}

func banner() {
	banner := `
                                                  __ 
 ____  _____ _____    __            _   _____    |  |
|    \|   | |   __|  |  |   ___ ___| |_|  |  |___|  |
|  |  | | | |__   |  |  |__| . | . | '_|  |  | . |__|
|____/|_|___|_____|  |_____|___|___|_,_|_____|  _|__|
                                             |_|     
	
											 `
	fmt.Println(banner)
}

func main() {
	banner()
	var urlStr string
	var falseBool bool
	var trueBool bool
	var allBool bool
	flag.StringVar(&urlStr, "u", "example", "Search url")
	flag.BoolVar(&falseBool, "f", false, "FAIL")
	flag.BoolVar(&trueBool, "t", false, "FOUND")
	flag.BoolVar(&allBool, "a", false, "ALL")
	flag.Parse()

	// Check if allBool is true
	if allBool == true {
		falseBool = true
		trueBool = true
	}

	// Open the file
	filePath := "output.txt"
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Process each line
	for scanner.Scan() {
		line := scanner.Text()

		// Do something with the line
		lowerStr := strings.ToLower(line)
		done := urlStr + "." + lowerStr
		DNS(done, falseBool, trueBool)
	}

	// Check for any errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error while scanning file: %v\n", err)
		return
	}
}

func DNS(hostname string, falseBool bool, trueBool bool) {
	// Specify the hostname to check

	// Perform a DNS lookup for the A record
	ips, err := net.LookupIP(hostname)
	if err != nil {
		if falseBool == true {
			u := hostname
			f := " ... FAIL"
			r := len(u)
			w := len(f)
			for i := 0; i < 20; i++ {
				d := i + r + w
				if d == 30 {
					g := u + strings.Repeat(" ", i) + f
					falseLogger.Printf("\x1b[31m%s\x1b[0m\n", g)
				}
			}
		}
		return
	}

	// Check if any IP addresses are found
	if len(ips) > 0 {
		if trueBool == true {
			u := hostname
			f := " ... FOUND"
			r := len(u)
			w := len(f)
			for i := 0; i < 20; i++ {
				d := i + r + w
				if d == 31 {
					g := u + strings.Repeat(" ", i) + f
					trueLogger.Printf("\x1b[32m%s\x1b[0m\n", g)
				}
			}
		}

	}

}
