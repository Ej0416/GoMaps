package main

import (
	"fmt"
)

// func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
// 	if len(names) != len(phoneNumbers) {
// 		return nil, errors.New("invalid sizes")
// 	}

// 	userMap := make(map[string]user)

// 	for i := range names {
// 		userStruct := user{name: names[i], phoneNumber: phoneNumbers[i]}
// 		userMap[names[i]] = userStruct;
// 	}

// 	return userMap, nil
// }

// type user struct {
// 	name        string
// 	phoneNumber int
// }

// func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
// 	elem, ok := users[name]
// 	if !ok {
// 		return false, errors.New("not found")
// 	}

// 	if !elem.scheduledForDeletion {
// 		return false, nil
// 	}

// 	delete(users,name)
// 	return true, nil
// }

// type user struct {
// 	name                 string
// 	number               int
// 	scheduledForDeletion bool
// }

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	for _,m := range messagedUsers {
		if _, ok := validUsers[m]; ok{
			validUsers[m] += 1
		}
	}
}

func main() {
	fmt.Println("app start")
}
