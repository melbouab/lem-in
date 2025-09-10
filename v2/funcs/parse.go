package helpers

import (
	"fmt"
	"strconv"
	"strings"
)

// it Splits the input into 3 parts:
// Ants Number : Checks if its valid using the function GetAnts.
// Rooms Slice : Checks if all the room lines are valid using the function GetRooms.
// Links Slice : Checks if all the links are valid using the function GetLinks.
func InputSplit(lines []string) error {
	for i := 0; i < len(lines); i++ {
		for strings.Contains(lines[i], "  ") {
			lines[i] = strings.ReplaceAll(lines[i], "  ", " ")
		}
		lines[i] = strings.TrimSpace(lines[i])
		if lines[i] == "" {
			lines = append(lines[:i], lines[i+1:]...)
			i--
		}
	}

	if len(lines) == 0 {
		return fmt.Errorf("ERROR: input empty")
	}

	err := GetAnts(lines[0])
	if err != nil {
		return err
	}
	lines = lines[1:]
	rooms := []string{}
	start := 0
	end := 0
	for i, line := range lines {
		if strings.HasPrefix(line, "#") {
			if line == "##start" {
				rooms = append(rooms, line)
				start++
			}
			if line == "##end" {
				rooms = append(rooms, line)
				end++
			}
			continue
		}
		if strings.Contains(line, "-") {
			lines = lines[i:]
			break
		}
		rooms = append(rooms, line)
	}
	err1 := GetRooms(rooms)
	err2 := GetLinks(lines)
	if err1 != nil {
		return err1
	}
	if err2 != nil {
		return err2
	}
	if start != 1 || end != 1 || rooms[len(rooms)-1] == "##end" || rooms[len(rooms)-1] == "##start" {
		return fmt.Errorf("ERROR: invalid input")
	}
	return nil
}

// Gets ants number and return error in case of invalid or non existing number.
func GetAnts(a string) error {
	a = strings.TrimSpace(a)
	ants, err := strconv.Atoi(a)
	if err != nil {
		return fmt.Errorf("ERROR: invalid ants number < %s >", a)
	}
	if ants < 1 {
		return fmt.Errorf("ERROR: ants number cannot be less than 1 < %s >", a)
	}
	Ants = ants
	return nil
}

// Rooms Part:

// IsDigit Checks if the stirng passsed is a number.
func IsNumber(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Checks if the room is Valid.
var existCoord = make(map[string]bool)

func ValidRoom(r string, room []string) error {
	if len(room) != 3 {
		return fmt.Errorf("ERROR: invalid room < %s >", r)
	}
	if !IsNumber(room[1]) || !IsNumber(room[2]) {
		return fmt.Errorf("ERROR: invalid room coordinates < %s >", r)
	}
	if existCoord[room[1]+" "+room[2]] {
		return fmt.Errorf("ERROR: room coordinates must be unique (%s , %s)", room[1], room[2])
	}
	existCoord[room[1]+" "+room[2]] = true
	if strings.HasPrefix(room[0], "L") {
		return fmt.Errorf("ERROR: a room should not start with letter < L >")
	}
	_, exist := Graph[room[0]]
	if exist {
		return fmt.Errorf("ERROR: room names must be unique")
	}
	return nil
}

// Fills the Graph with rooms.
func GetRooms(rooms []string) error {
	main := ""
	for _, room := range rooms {
		if room == "##start" {
			main = "s"
			continue
		}
		if room == "##end" {
			main = "e"
			continue
		}
		elems := strings.Fields(room)
		err := ValidRoom(room, elems)
		if err != nil {
			return err
		}
		newroom := &Room{
			Name: elems[0],
		}
		if main == "s" {
			Start = elems[0]
		}
		if main == "e" {
			End = elems[0]
		}
		Graph[elems[0]] = newroom
		main = ""
	}
	return nil
}

// Rooms End

// Links Part:

// Checks if the links is valid.
func ValidLink(l string, link []string) error {
	if len(link) != 2 {
		return fmt.Errorf("ERROR: invalid link < %s >", l)
	}
	if link[0] == link[1] {
		return fmt.Errorf("ERROR: invalid link < %s >", l)
	}
	_, exits1 := Graph[link[0]]
	_, exits2 := Graph[link[1]]
	if !exits1 {
		return fmt.Errorf("ERROR: room linked does not exist < %s >", link[0])
	}
	if !exits2 {
		return fmt.Errorf("ERROR: room linked does not exist < %s >", link[1])
	}
	return nil
}

// Links rooms in the graph with each others.
var linkexists = make(map[string]bool)

func GetLinks(links []string) error {
	for _, link := range links {
		if strings.HasPrefix(link, "#") {
			continue
		}
		elems := strings.Split(link, "-")
		err := ValidLink(link, elems)
		if err != nil {
			return err
		}
		l1, l2 := elems[0]+"-"+elems[1], elems[1]+"-"+elems[0]
		if linkexists[l1] || linkexists[l2] {
			continue
		}
		Graph[elems[0]].Neighbors = append(Graph[elems[0]].Neighbors, Graph[elems[1]])
		Graph[elems[1]].Neighbors = append(Graph[elems[1]].Neighbors, Graph[elems[0]])
		linkexists[l1], linkexists[l2] = true, true
	}
	return nil
}
