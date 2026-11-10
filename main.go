package main

func main() {}
func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	totalCost := 0
	costs := [3]int{}
	for i := 0; i < len(messages); i++ {
		messageCost := len(messages[i])
		totalCost += messageCost
		costs[i] = totalCost
	}
	return messages, costs
}
