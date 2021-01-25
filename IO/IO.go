package IO

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func WriteTo1(w io.Writer, lines []string) error {
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func WriteTo2(w io.Writer, lines []string) (n int64, err error) { // 반환형에 명명된 결과 인자를 지정 하여 n,err를 선언하지 않아도 반환할수 있다.
	for _, line := range lines {
		var nw int
		nw, err = fmt.Fprintln(w, line)
		n += int64(nw)
		if err != nil {
			return
		}
	}
	return // 명명된 결과 인자를 사용할 경우 반환 인자를 생략할 수 있다. 생략시 기본값을 초기화됨
	// 명명된 결과 인자는 코드를 읽기 어렵게 만들기 때문에
}

func WriteTo3(w io.Writer, lines ...string) error { // lines를 []string으로 받을 수 있지만 ...로 받을 수 있다.
	for _, line := range lines {
		if _, err := fmt.Fprintln(w, line); err != nil {
			return err
		}
	}
	return nil
}

func ReadFrom(r io.Reader, lines *[]string) error {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		*lines = append(*lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func ExampleWriteTo() {
	lines := []string{
		"bill@mail.com",
		"tom@mail.com",
		"jane@mail.com",
	}
	if err := WriteTo1(os.Stdout, lines); err != nil {
		fmt.Println(err)
	}
	if err := WriteTo3(os.Stdout, lines...); err != nil {
		fmt.Println(err)
	}
}

func ExampleReadFrom() {
	r := strings.NewReader("bill\ntome\njane\n")
	var lines []string
	if err := ReadFrom(r, &lines); err != nil {
		fmt.Println(err)
	}
	fmt.Println(lines)
}
