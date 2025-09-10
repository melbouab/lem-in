package main

import (
	"fmt"
	"os"
	"strings"

	helpers "lemin/funcs"
)

func main() {
	if len(os.Args) != 2 || !strings.HasSuffix(os.Args[1], ".txt") {
		fmt.Println("Usage: go run . fileName.txt")
		return
	}
	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	lines := strings.Split(string(content), "\n")
	err = helpers.InputSplit(lines)
	if err != nil {
		fmt.Println(err)
		return
	}
	allGroups := helpers.AllGroups()
	if len(allGroups) == 0 {
		fmt.Println("ERROR: no path is found")
		return
	}

	fmt.Println(string(content) + "\n")
	helpers.MoveAnts(allGroups)
}
