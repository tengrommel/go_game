package main

import (
	"context"
	"fmt"
)

// user is the type of value to store in the context.
type user struct {
	name string
}

// userKey is the type of value to use for the key.
// The key is type specific and only values of the same type
type userKey int

func main() {
	// Create a value of type user.
	u := user{
		name: "Bill",
	}

	// Declare a key with the value of zero of type userKey.
	const uk userKey = 0

	// Store the pointer to user value inside the context with a value of zero of type userKey.
	ctx := context.WithValue(context.Background(), uk, &u)

	// Retrieve that user pointer back by user the same key type value.
	if u, ok := ctx.Value(uk).(*user);ok{
		fmt.Println("User", u.name)
	}

	// Attempt to retrieve the value again using the same value but of different type.
	if _, ok:= ctx.Value(0).(*user);!ok{
		fmt.Println("User Not Found")
	}

}