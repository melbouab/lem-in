package lemin

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name string
	X, Y int
}

type Colony struct {
	NumAnts    int
	Rooms      map[string]*Room
	Links      map[string][]string
	Start      string
	End        string
	Valid      map[string]int
	ValidCord  map[[2]int]int
	InputLines []string
}

type Path struct {
	Rooms  []string
	Length int
}

func ParseInput() (*Colony, error) {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <filename>")
		os.Exit(1)
	}
	filename := os.Args[1]
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	colony := &Colony{
		Rooms:     make(map[string]*Room),
		Links:     make(map[string][]string),
		Valid:     map[string]int{},
		ValidCord: map[[2]int]int{},
	}

	colony.InputLines = strings.Split(string(file), "\n")

	if len(colony.InputLines) == 0 {
		return nil, fmt.Errorf("empty file")
	}
	for _, v := range colony.InputLines {
		colony.Valid[v]++
		if colony.Valid[v] > 1 {
			return nil, fmt.Errorf("invalid format duplicate Romms")
		}
	}
	colony.Valid = make(map[string]int)
	numAnts, err := strconv.Atoi(colony.InputLines[0])
	if err != nil || numAnts <= 0 {
		return nil, fmt.Errorf("invalid number of ants")
	}
	colony.NumAnts = numAnts

	i := 1
	var nextIsStart, nextIsEnd bool

	for i < len(colony.InputLines) {
		line := strings.TrimSpace(colony.InputLines[i])

		if line == "" || (strings.HasPrefix(line, "#") && line != "##start" && line != "##end") {
			i++
			continue
		}

		if line == "##start" {
			nextIsStart = true
			i++
			continue
		}
		if line == "##end" {
			nextIsEnd = true
			i++
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			if len(parts) != 2 {
				return nil, fmt.Errorf("invalid link format")
			}

			room1, room2 := parts[0], parts[1]
			if colony.Rooms[room1] == nil || colony.Rooms[room2] == nil {
				return nil, fmt.Errorf("link to unknown room")
			}

			colony.Links[room1] = append(colony.Links[room1], room2)
			colony.Links[room2] = append(colony.Links[room2], room1)
		} else {
			parts := strings.Fields(line)
			if len(parts) != 3 {
				return nil, fmt.Errorf("invalid room format")
			}

			name := parts[0]
			if strings.HasPrefix(name, "L") || strings.HasPrefix(name, "#") || strings.Contains(name, " ") {
				return nil, fmt.Errorf("invalid room name")
			}

			x, err1 := strconv.Atoi(parts[1])
			y, err2 := strconv.Atoi(parts[2])
			if err1 != nil || err2 != nil {
				return nil, fmt.Errorf("invalid coordinates")
			}

			if colony.Rooms[name] != nil {
				return nil, fmt.Errorf("duplicate room")
			}

			colony.Rooms[name] = &Room{Name: name, X: x, Y: y}

			if nextIsStart {
				colony.Start = name
				nextIsStart = false
			}
			if nextIsEnd {
				colony.End = name
				nextIsEnd = false
			}
		}
		i++
	}
	for _, v := range colony.Rooms {
		colony.Valid[v.Name]++
		if colony.Valid[v.Name] > 1 {
			return nil, fmt.Errorf("invalid format duplicate Rome")
		}
		colony.ValidCord[[2]int{v.X, v.Y}]++
		if colony.ValidCord[[2]int{v.X, v.Y}] > 1 {
			return nil, fmt.Errorf("invalid duplicat Cordonit")
		}
	}

	if colony.Start == "" || colony.End == "" {
		return nil, fmt.Errorf("missing start or end room")
	}

	return colony, nil
}
