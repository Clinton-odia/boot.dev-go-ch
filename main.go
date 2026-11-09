package main

func countConnections(groupSize int) int {
	var connection int

	for i := 0; i < groupSize; i++ {
		connection = connection + i
	}

	return connection
}
