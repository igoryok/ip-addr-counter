package main

import (
	"testing"
)

func TestCountUniqueIPs(t *testing.T) {

	// Call the function and verify the result
	count, err := CountUniqueIPs("ip_addresses.txt")
	if err != nil {
		t.Fatalf("Error counting unique IPs: %v", err)
	}

	expectedCount := 10
	if count != expectedCount {
		t.Errorf("Expected %d unique IPs, got %d", expectedCount, count)
	}
}
