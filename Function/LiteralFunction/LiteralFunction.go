package LiteralFunction

import (
	"fmt"
)

func add(a, b int) int { // 기본 함수
	return a + b
}

// func(a, b int) int { // 리터럴 함수
// 	return a + b
// }

func printHello() { // 기본 함수
	fmt.Println("Hello")
}

func Example_functionLiteral() { // 리터럴 함수 사용
	func() {
		fmt.Println("Hello")
	}()
}

func Example_functionLiteralVal() { // 리터럴 함수를 변수에 담에서 호
	printBye := func() {
		fmt.Println("Bye")
	}
	printBye()
}
