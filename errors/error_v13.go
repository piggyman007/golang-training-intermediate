package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var err error
	originalError := &os.PathError{Op: "read", Path: "/a/b", Err: errors.New("read error")}

	err = fmt.Errorf("context for error: %v", originalError)
	fmt.Printf("error type: %T\n", err)
	fmt.Printf("error value: %v\n", err)
	fmt.Println()

	// %w - wrap
	err = fmt.Errorf("wrapping for error: %w", originalError)
	fmt.Printf("error type: %T\n", err) // print *fmt.wrapError
	fmt.Printf("error value: %v\n", err)
	fmt.Println()

	err = errors.Unwrap(err)
	fmt.Printf("error type: %T\n", err) // print *os.PathError
	fmt.Printf("error value: %v\n", err)
	fmt.Println()

	err = errors.Unwrap(err)
	fmt.Printf("error type: %T\n", err) // print *errors.errorString
	fmt.Printf("error value: %v\n", err)
	fmt.Println()

	err = errors.Unwrap(err)
	fmt.Printf("error type: %T\n", err) // print nil
	fmt.Printf("error value: %v\n", err)
	fmt.Println()

	err = errors.Unwrap(err)
	fmt.Printf("error type: %T\n", err) // print nil
	fmt.Printf("error value: %v\n", err)
	fmt.Println()
}
