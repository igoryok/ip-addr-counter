package main

import (
	"log"
	"time"
)

func main() {
	start := time.Now()

	filePath := "ip_addresses.txt"
	count, err := CountUniqueIPs(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	log.Printf("Number of unique IPs - %d", count)

	elapsed := time.Since(start)
	log.Printf("Elapsed Time in milliseconds: %d", elapsed.Milliseconds())
}
