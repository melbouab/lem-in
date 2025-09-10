package helpers

import (
	"slices"
)

type Room struct {
	Name      string
	Neighbors []*Room
}

var (
	Ants            int
	Start           string
	End             string
	directPathAdded bool
)

var Graph = make(map[string]*Room)

func AllGroups() [][][]string {
	basePaths := [][]string{}
	for _, firstChild := range Graph[Start].Neighbors {
		basePaths = append(basePaths, bfs(firstChild.Name, nil))
	}
	groups := [][][]string{}
	for _, basePath := range basePaths {

		forbidden := make(map[string]bool)
		for _, room := range basePath[:len(basePath)-1] {
			forbidden[room] = true
		}
		group := [][]string{basePath}
		for {
			path := bfs(Start, forbidden)
			if path == nil {
				break
			}
			for _, room := range path[:len(path)-1] {
				forbidden[room] = true
			}
			group = append(group, path)
		}
		groups = append(groups, group)
	}
	return groups
}

func bfs(start string, forbidden map[string]bool) []string {
	visited := make(map[string]bool)
	visited[Start], visited[start] = true, true
	parents := make(map[string]string)
	queue := []string{start}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr == End {
			path := gPath(start, parents)

			if forbidden != nil && len(path) == 2 {
				if directPathAdded {
					continue
				}
				directPathAdded = true
			}

			if forbidden != nil {
				return path[1:]
			}

			return path
		}

		for _, neighbor := range Graph[curr].Neighbors {
			if visited[neighbor.Name] || (forbidden != nil && forbidden[neighbor.Name]) {
				continue
			}
			visited[neighbor.Name] = true
			parents[neighbor.Name] = curr
			queue = append(queue, neighbor.Name)
		}
	}
	return nil
}

func gPath(start string, parents map[string]string) []string {
	curr := End
	path := []string{}
	for curr != start {
		path = append(path, curr)
		curr = parents[curr]
	}
	path = append(path, start)
	slices.Reverse(path)
	return path
}
