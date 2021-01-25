package Error

import (
	"errors"
	"os"
	"preview/IO"
)

func ExampleError() (*int, error) {
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
	}
	if _, err := IO.WriteTo2(os.Stdout, lines); err != nil {
		return nil, err
	}
	// error 생성하는 방법
	return nil, errors.New("Error occured")
}
