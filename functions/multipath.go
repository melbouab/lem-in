package lemin

func FindMultiplePaths(colony *Colony) []Path {
	var allPaths []Path

	// Find shortest path first using BFS
	shortestPath := FindShortestPath(colony)
	if shortestPath == nil {
		return nil
	}
	allPaths = append(allPaths, *shortestPath)

	// Try to find additional non-overlapping paths
	usedRooms := make(map[string]bool)
	for i := 1; i < len(shortestPath.Rooms)-1; i++ {
		usedRooms[shortestPath.Rooms[i]] = true
	}

	// Look for alternative paths that don't share intermediate rooms
	for attempts := 0; attempts < 3; attempts++ {
		altPath := FindAlternativePath(colony, usedRooms)
		if altPath != nil && altPath.Length <= shortestPath.Length+2 {
			allPaths = append(allPaths, *altPath)
			// Mark rooms in this path as used
			for i := 1; i < len(altPath.Rooms)-1; i++ {
				usedRooms[altPath.Rooms[i]] = true
			}
		}
	}

	return allPaths
}
