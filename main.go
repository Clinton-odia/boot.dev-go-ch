package main

func countReports(numSentCh chan int) int {
	numReport := 0
	for {
		reportSent, ok := <-numSentCh
		if !ok {
			break
		}
		numReport += reportSent

	}
	return numReport
}

// don't touch below this line

func sendReports(numBatches int, ch chan int) {
	for i := 0; i < numBatches; i++ {
		numReports := i*23 + 32%17
		ch <- numReports
	}
	close(ch)
}
