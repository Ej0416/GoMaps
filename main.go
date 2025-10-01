package main

import (
	"fmt"
	"strings"
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

// func updateCounts(messagedUsers []string, validUsers map[string]int) {
// 	for _,m := range messagedUsers {
// 		if _, ok := validUsers[m]; ok{
// 			validUsers[m] += 1
// 		}
// 	}
// }

// func getNameCounts(names []string) map[rune]map[string]int {
// 	nameCounts := make(map[rune]map[string]int)

// 	for _, n := range names {
// 		fc := []rune(n)[0];
// 		// fc := unicode.ToLower([]rune(n)[0]); // normalize version
// 		if _,ok := nameCounts[fc]; !ok{
// 			nameCounts[fc] = make(map[string]int)
// 		}
// 		nameCounts[fc][n]++
// 	}

// 	return nameCounts
// }

func countDistinctWords(messages []string) int {
	distinctWords := make(map[string]int);


	for _,message := range messages{
		words := strings.FieldsSeq(message)
		for word := range words {
			distinctWords[strings.ToLower(word)]++
		}
	}

	return len(distinctWords)
}


func main() {
	fmt.Println("app start")
	messages := []string{"Hello world", "hello there", "General Kenobi"}
	fmt.Println(countDistinctWords(messages))
}
