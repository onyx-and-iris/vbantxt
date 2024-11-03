package main

import (
	"flag"
	"slices"
)

func flagsPassed(flags []string) bool {
	found := false
	flag.Visit(func(f *flag.Flag) {
		if slices.Contains(flags, f.Name) {
			found = true
			return
		}
	})
	return found
}

func indexOf[T comparable](collection []T, e T) int {
	for i, x := range collection {
		if x == e {
			return i
		}
	}
	return -1
}
