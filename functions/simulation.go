package lemin

import (
	"fmt"
	"sort"
	"strings"
)

func SimulateMultiPathMovement(colony *Colony, paths []Path) {
	if len(paths) == 0 {
		return
	}

	// Distribute ants across paths optimally
	antPaths := DistributeAntsOptimally(colony.NumAnts, paths)

	// Initialize ant positions
	type AntState struct {
		PathIndex int
		Position  int // Position in the path
	}

	ants := make([]AntState, colony.NumAnts)
	antID := 0

	for pathIdx, count := range antPaths {
		for i := 0; i < count; i++ {
			ants[antID] = AntState{PathIndex: pathIdx, Position: 0}
			antID++
		}
	}

	for {
		var moves []string
		activeAnts := 0

		// Check how many ants are still active
		for i := 0; i < colony.NumAnts; i++ {
			ant := &ants[i]
			path := paths[ant.PathIndex]
			if ant.Position < len(path.Rooms)-1 {
				activeAnts++
			}
		}

		if activeAnts == 0 {
			break
		}

		// Try to move ants
		nextPositions := make([]int, colony.NumAnts)
		for i := 0; i < colony.NumAnts; i++ {
			nextPositions[i] = ants[i].Position
		}

		for i := colony.NumAnts - 1; i >= 0; i-- {
			ant := &ants[i]
			path := paths[ant.PathIndex]

			if ant.Position >= len(path.Rooms)-1 {
				continue
			}

			nextPos := ant.Position + 1
			nextRoom := path.Rooms[nextPos]
			canMove := true

			// Check for conflicts with other ants
			if nextRoom != colony.End {
				for j := 0; j < colony.NumAnts; j++ {
					if i != j {
						otherAnt := &ants[j]
						otherPath := paths[otherAnt.PathIndex]
						if nextPositions[j] < len(otherPath.Rooms) &&
							otherPath.Rooms[nextPositions[j]] == nextRoom {
							canMove = false
							break
						}
					}
				}
			}

			if canMove {
				nextPositions[i] = nextPos
				moves = append(moves, fmt.Sprintf("L%d-%s", i+1, nextRoom))
			}
		}

		// Update positions
		for i := 0; i < colony.NumAnts; i++ {
			ants[i].Position = nextPositions[i]
		}

		if len(moves) > 0 {
			// Sort moves by ant ID for consistent output
			sort.Strings(moves)
			fmt.Println(strings.Join(moves, " "))
		}
	}
}

func DistributeAntsOptimally(numAnts int, paths []Path) []int {
	distribution := make([]int, len(paths))

	// Simple greedy distribution
	for i := 0; i < numAnts; i++ {
		bestIdx := 0
		bestCost := distribution[0] + paths[0].Length

		for j := 1; j < len(paths); j++ {
			cost := distribution[j] + paths[j].Length
			if cost < bestCost {
				bestCost = cost
				bestIdx = j
			}
		}

		distribution[bestIdx]++
	}

	return distribution
}
