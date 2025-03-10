// Package vbantxt provides utilities for handling VBAN text errors.
package vbantxt

// Error is used to define sentinel errors.
type Error string

// Error implements the error interface.
func (r Error) Error() string {
	return string(r)
}
