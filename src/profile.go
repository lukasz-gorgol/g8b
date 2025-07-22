//go:build profile
// +build profile

package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

// Initialize profiling when this file is built with the profile tag
func init() {
	// Create CPU profile file
	f, err := os.Create("cpu.pprof")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not create CPU profile: %v\n", err)
		return
	}

	// Start CPU profiling
	if err := pprof.StartCPUProfile(f); err != nil {
		fmt.Fprintf(os.Stderr, "Could not start CPU profile: %v\n", err)
		return
	}

	// Register cleanup function to stop profiling when the program exits
	fmt.Println("CPU profiling enabled. Profile will be written to cpu.pprof")

	// Setup deferred stop of CPU profiling
	go func() {
		defer pprof.StopCPUProfile()
		defer f.Close()

		// Wait for program to finish
		<-make(chan struct{})
	}()
}
