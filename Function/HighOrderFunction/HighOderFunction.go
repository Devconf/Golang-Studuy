package HighOrderFunction

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func ReadFrom(r io.Reader, f func(line string)) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ExampleReadFrom_Print() {
	r := strings.NewReader("bill\ntome\njane\n")
	err := ReadFrom(r, func(line string) {
		fmt.Println("(", line, ")")
	})
	if err != nil {
		fmt.Println(err)
	}
}

// 클로저
func ExampleReadFrom_append() {
	r := strings.NewReader("bill\ntome\njane\n")
	var lines []string //클로저는 외부에서 선언한 변수를 함수 리터럴 내에서 마음대로 접근할 수 있는 코드를 의미한다.
	err := ReadFrom(r, func(line string) {
		lines = append(lines, line)
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}
