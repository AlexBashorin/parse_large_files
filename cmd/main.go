package main

import (
	"fmt"

	"ecwid/internal/parse"
)

func main() {
	filename := "../large_ip_file.txt"
	uniqueCount, err := parse.CountUniqueIPs(filename)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("Number of unique IP addresses: %d\n", uniqueCount)
}
