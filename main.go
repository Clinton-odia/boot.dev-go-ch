package main

func maxMessages(thresh int) int {
	// ?
	var totalCost int
	var messageCount int
	for i := 0; ; i++ {
		messageCost := 100 + i
		if totalCost+messageCost > thresh {
			break
		}
		messageCount++
		totalCost += messageCost

	}
	return messageCount
}
