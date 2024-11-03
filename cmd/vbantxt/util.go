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
