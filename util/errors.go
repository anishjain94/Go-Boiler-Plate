package util

import (
	"fmt"
)

func AssertError(err error, statusCode int, errorMsg string) {
	if err != nil {

		errorMsg := fmt.Sprint(statusCode) + ":" + errorMsg
		panic(errorMsg)
	}
}

func ErrorIf(condition bool, statusCode int, errorMsg string) {
	if condition {

		errorMsg := fmt.Sprint(statusCode) + ":" + errorMsg
		panic(errorMsg)
	}
}

func Assert(condition bool, statusCode int, errorMsg string) {
	if !condition {
		errorMsg := fmt.Sprint(statusCode) + ":" + errorMsg
		panic(errorMsg)
	}
}
