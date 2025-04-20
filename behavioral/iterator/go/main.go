package main

import "fmt"

func main() {
	user1 := &User{"a", 30}
	user2 := &User{"b", 20}

	users := &UserCollection{[]*User{user1, user2}}

	iter := users.createIterator()
	for iter.hasNext() {
		user := iter.getNext()
		fmt.Printf("User is %+v\n", user)
	}
}
