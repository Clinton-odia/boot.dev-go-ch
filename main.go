package main

//import "fmt"

func isValidPassword(password string) bool {
	if len(password) >= 5 && len(password) <= 12 {
		var isCap bool
		var isDigit bool
		for _, v := range password {

			if v >= 'A' && v <= 'Z' {
				isCap = true
			}
			if v >= '0' && v <= '9' {
				isDigit = true
			}

		}
		return isCap && isDigit

	}
	return false
}

// func main() {
// 	isValidPassword("Pass123")
// }
