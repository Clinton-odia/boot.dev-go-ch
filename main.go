package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	users := map[string]user{}
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	} else {
		for i := 0; i < len(names); i++ {
			value := names[i]
			users[value] = user{name: names[i], phoneNumber: phoneNumbers[i]}

		}
	}

	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}
