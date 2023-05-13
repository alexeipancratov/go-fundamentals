package errors

import (
	"errors"
	"fmt"
)

func Demo() {
	result, err := divide2(100, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(result)
}

// For cases when we want to handle error states ourselves
func divide1(l, r int) (int, error) {
	if r == 0 {
		return 0, errors.New("can't divide by zero")
	}
	return l / r, nil
}

// For cases when any runtime error occurs down the stack and we want to recover from it
func divide2(l, r int) (result int, err error) {
	defer func() {
		if msg := recover(); msg != nil {
			result = 0
			err = fmt.Errorf("%v", msg)
		}
	}()
	return l / r, nil
}
