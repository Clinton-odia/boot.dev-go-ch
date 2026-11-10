package main

type cost struct {
	day   int
	value float64
}

func getDayCosts(costs []cost, day int) []float64 {
	dayCosts := make([]float64, 0)
	for _, c := range costs {
		if c.day == day {
			dayCosts = append(dayCosts, c.value)
		}
	}
	return dayCosts
}
