package main

import (
	"errors"
	"fmt"
	"os"
)

func insertInfo() error {
	return &os.PathError{Op: "read", Path: "/a/b", Err: errors.New("read error")}
}

func main() {
	originalError := insertInfo()

	// // old version
	// if e, ok := originalError.(*os.PathError); ok {
	// 	fmt.Println("path error => ", e)
	// }

	var pErr *os.PathError
	if errors.As(originalError, &pErr) {
		fmt.Println("path error => ", pErr)
	}
}
