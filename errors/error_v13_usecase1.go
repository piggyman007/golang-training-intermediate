package main

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
)

func insertInfo() error {
	return fmt.Errorf("error => %w", sql.ErrTxDone)
}

func main() {
	var err error

	err = insertInfo()

	// steps
	// 1. use %w to wrap
	// 2. check err by using errors.Is() function
	if errors.Is(err, sql.ErrTxDone) {
		fmt.Println("err => ", err)
	}

	if err != nil {
		// check nil should be the same as version < 1.13
		return
	}

	if err == io.EOF {
		// io.EOF should not be wrap
		return
	}
}
