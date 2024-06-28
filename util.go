package main

import "flag"

// indexOf returns the index of an element in an array
func indexOf[T comparable](collection []T, e T) int {
	for i, x := range collection {
		if x == e {
			return i
		}
	}
	return -1
}

func isFlagPassed(name string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == name {
			found = true
		}
	})
	return found
}
