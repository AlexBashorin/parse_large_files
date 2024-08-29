package main

import (
	"ecwid/internal/parse"
	"fmt"
)

func main() {
	filename := "../large_ip_file.txt"
	uniqueCount, pathToUniqIPFile, err := parse.CountUniqueIPs(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Number of unique IP: %d\n", uniqueCount)
	fmt.Printf("Path to the file with unique IPs: %s\n", pathToUniqIPFile)
}
