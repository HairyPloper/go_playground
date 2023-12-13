// Package main is the entry point for the Congestion Calculator Manager application.

package main

import "congestion-calculator-manager/app/server"

// main is the entry point for the Congestion Calculator Manager application.
// It invokes the StartServer function from the server package to initialize and start the server.
func main() {
	server.StartServer()
}
