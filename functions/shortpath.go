package lemin

func FindShortestPath(colony *Colony) *Path {
	type QueueItem struct {
		room   string
		parent *QueueItem
		depth  int
	}

	queue := []*QueueItem{{room: colony.Start, parent: nil, depth: 0}}
	visited := map[string]bool{colony.Start: true}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.room == colony.End {
			path := make([]string, current.depth+1)
			for i, item := current.depth, current; item != nil; i, item = i-1, item.parent {
				path[i] = item.room
			}
			return &Path{Rooms: path, Length: current.depth}
		}

		for _, neighbor := range colony.Links[current.room] {
			if !visited[neighbor] {
				visited[neighbor] = true
				queue = append(queue, &QueueItem{
					room:   neighbor,
					parent: current,
					depth:  current.depth + 1,
				})
			}
		}
	}
	return nil
}
