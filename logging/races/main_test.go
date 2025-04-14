package main

import (
	"sync"
	"testing"

	"k8s.io/klog/v2"
)

func TestLogSomethingConcurrent(t *testing.T) {
	// Initialize klog
	klog.InitFlags(nil)

	var wg sync.WaitGroup
	numGoroutines := 1009

	// Run the function concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			logSomething()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
