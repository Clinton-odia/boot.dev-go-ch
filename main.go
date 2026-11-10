package main

func getMessageCosts(messages []string) []float64 {
	messageCost := make([]float64, 0, len(messages))
	for _, message := range messages {
		messageCost = append(messageCost, float64(len(message))*0.01)
	}
	return messageCost
}
