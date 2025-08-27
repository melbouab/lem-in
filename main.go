
package main

import (
	"fmt"
	"os"

	lemin "lemin/functions"
)

func main() {
	colony, err := lemin.ParseInput()
	if err != nil {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	// Find multiple paths and select optimal combination
	paths := lemin.FindMultiplePaths(colony)
	if len(paths) == 0 {
		fmt.Println("ERROR: invalid data format")
		os.Exit(1)
	}

	// Print the input
	for _, line := range colony.InputLines {
		fmt.Println(line)
	}
	fmt.Println()

	// Simulate ant movement with multiple paths
	lemin.SimulateMultiPathMovement(colony, paths)
}