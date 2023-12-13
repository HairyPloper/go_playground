// Package main is the entry point for the Congestion Calculator Manager application.

package main

import (
	"congestion-calculator-manager/app/client"
)

// main is the entry point for the Congestion Calculator Manager application.
// It invokes the StartClient function from the client package to initiate the client-side operations.
func main() {
	client.StartClient()
}
