package helpers

import (
	"fmt"
	"sort"
	"strings"
)

type DistributeInfo struct {
	ID        int
	AntNumber int
	PathLen   int
}

// AddExpand appends an ant's moves into allTurns at the correct starting turn.
// Expands allTurns if necessary.
func AddExpand(allTurns [][]string, start int, moves []string) [][]string {
	for len(allTurns) < start+len(moves) {
		allTurns = append(allTurns, nil)
	}

	for i, move := range moves {
		allTurns[start+i] = append(allTurns[start+i], move)
	}

	return allTurns
}

// printTurns outputs the turns, one per line.
func printTurns(turns [][]string) {
	for _, turn := range turns {
		fmt.Println(strings.Join(turn, " "))
	}
}

// MoveAnts simulates ants moving through paths and records their moves.
func MoveAnts(allGroups [][][]string) {
	bestGroup, assignedAnts := SelectGroup(allGroups)
	allTurns := [][]string{}

	turn := 0
	for len(assignedAnts[0]) > 0 {
		for i, ants := range assignedAnts {
			if len(ants) == 0 {
				continue
			}

			currAnt := ants[0]
			moves := make([]string, len(bestGroup[i]))
			for j, room := range bestGroup[i] {
				moves[j] = fmt.Sprintf("L%d-%s", currAnt, room)
			}

			allTurns = AddExpand(allTurns, turn, moves)
			assignedAnts[i] = ants[1:]
		}
		turn++
	}

	printTurns(allTurns)
}

// AssignAnts distributes ants evenly across paths.
func AssignAnts(paths [][]string) [][]int {
	n := len(paths)
	assign := make([][]int, n)

	longest := len(paths[n-1])
	antID, antsLeft := 1, Ants

	for i, path := range paths {
		diff := longest - len(path)
		for diff > 0 && antsLeft > 0 {
			assign[i] = append(assign[i], antID)
			antID++
			antsLeft--
			diff--
		}
	}

	for antsLeft > 0 {
		for i := 0; i < n && antsLeft > 0; i++ {
			assign[i] = append(assign[i], antID)
			antID++
			antsLeft--
		}
	}

	return assign
}

// FilterGroups sorts the subgroups in each group by their length.
func FilterGroups(groups [][][]string) [][][]string {
	filtered := make([][][]string, len(groups))

	for i, group := range groups {
		sort.Slice(group, func(a, b int) bool {
			return len(group[a]) < len(group[b])
		})
		filtered[i] = group
	}

	return filtered
}

// Selects the group with the minimum number of turns.
func SelectGroup(groups [][][]string) ([][]string, [][]int) {
	groups = FilterGroups(groups)

	var BestGroupID int
	var best [][]string
	var cur [][]string
	for gI, group := range groups {

		distributeInfo := DistributeAntsAcrossPaths(Ants, group)
		cur = nil

		for i, info := range distributeInfo {
			if info.AntNumber > 0 {
				cur = append(cur, group[i])
			}
		}

		update := false

		// Rule 1: prefer more paths
		if len(cur) > len(best) && Ants >= len(cur) {
			update = true
		} else if len(cur) == len(best) {
			// Rule 2: prefer fewer total rooms
			bestRooms, curRooms := 0, 0
			for _, p := range best {
				bestRooms += len(p)
			}
			for _, p := range cur {
				curRooms += len(p)
			}

			if curRooms < bestRooms {
				update = true
			} else if curRooms == bestRooms {
				// Rule 3: prefer smaller smallest path length
				bestMin, curMin := 2147483648, 2147483648 // try to give the path the max possible look at the algo works next
				for _, p := range best {
					if len(p) < bestMin {
						bestMin = len(p)
					}
				}
				for _, p := range cur {
					if len(p) < curMin {
						curMin = len(p)
					}
				}
				if curMin < bestMin {
					update = true
				}
			}
		}

		if update {
			BestGroupID = gI
			best = make([][]string, len(cur))
			for i := range cur {
				best[i] = append([]string(nil), cur[i]...)
			}
		}
	}
	return groups[BestGroupID], AssignAnts(groups[BestGroupID])
}

// DistributeAntsAcrossPaths distributes the given number of ants across unique paths as evenly as possible.
func DistributeAntsAcrossPaths(antNumber int, group [][]string) []DistributeInfo {
	distributeInfo := make([]DistributeInfo, len(group))

	for i := range distributeInfo {
		distributeInfo[i].ID = i
		distributeInfo[i].PathLen = len(group[i]) - 2 // remove start room and end
	}

	antInRooms := 0
	for antNumber != 0 {
		for i := range distributeInfo {
			if distributeInfo[i].PathLen <= antInRooms {
				distributeInfo[i].AntNumber++
				antNumber--
				if antNumber == 0 {
					break
				}
			}
		}
		antInRooms++
	}
	return distributeInfo
}
