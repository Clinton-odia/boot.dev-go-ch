package main

import "fmt"

func main() {
	//monthlyBillIncrease(2, 89, 102)
	fmt.Printf("%d\n", monthlyBillIncrease(2, 89, 102))
}
func monthlyBillIncrease(costPerSend, numLastMonth, numThisMonth int) int {
	var lastMonthBill int
	var thisMonthBill int
	lastMonthBill = getBillForMonth(costPerSend, numLastMonth)
	thisMonthBill = getBillForMonth(costPerSend, numThisMonth)
	return thisMonthBill - lastMonthBill
}

func getBillForMonth(costPerSend, messagesSent int) int {
	bill := costPerSend * messagesSent
	return bill
}
