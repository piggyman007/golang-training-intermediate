package main

import (
	"fmt"
)

type BusinessError struct {
	Code    float32
	Message string
}

func (be BusinessError) Error() string {
	return fmt.Sprintf("%v (%v)", be.Message, be.Code)
}

func main() {
	err := GetErr()

	fmt.Println(err.Error())
}

func GetErr() error {
	bErr := BusinessError{
		Code:    500.12,
		Message: "XXX",
	}

	return bErr
}
