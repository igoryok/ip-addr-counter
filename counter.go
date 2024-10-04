package main

import (
	"bufio"
	"os"
	"strings"
)

// CountUniqueIPs reads a file with IP addresses and returns the number of unique IP addresses.
func CountUniqueIPs(filePath string) (int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	bitSet := NewBitSet(1 << 32) // Initialize BitSet with 2^32 size

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ip := strings.TrimSpace(scanner.Text())
		if ip == "" {
			continue
		}
		bitSet.Set(textToLong(ip))
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return bitSet.Count(), nil
}
