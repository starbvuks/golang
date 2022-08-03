package main

import "fmt"

func Mapping() {
	// map def
	emails := make(map[string]string)

	emails["User1"] = "user1@email.com"
	emails["User2"] = "user2@email.com"
	emails["DeleteMe"] = "trashdata"

	fmt.Printf("%v\n", emails)
	fmt.Printf("%v\n\n", emails["User2"])

	delete(emails, "DeleteMe")
	fmt.Printf("%v\n", emails)

	// compressed syntax
	cleanerEmails := map[string]string{"First": "1", "Second": "2"}
	fmt.Printf("%v\n\n", cleanerEmails)

	// looping through
	id := []int{002, 005, 007, 023, 043, 065, 077}
	for i, id := range id {
		fmt.Printf("%d - ID: %d\n", i, id)
	}

	// add ids
	sum := 0
	for _, id := range id {
		sum += id
	}
	fmt.Printf("sum: %d\n", sum)

	// mapping with range
	for c, v := range cleanerEmails {
		fmt.Printf("%s: %s\n", c, v)
	}

	for k := range emails {
		fmt.Printf("name: %s\n", k)
	}
}
