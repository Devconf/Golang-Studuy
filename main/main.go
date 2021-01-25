package main

import (
	"fmt"
	"preview/Function/Error"
	"preview/Function/HighOrderFunction"
	"preview/Function/LiteralFunction"
	"preview/Function/method"
	"preview/IO"
	"preview/Map"
	"preview/Set"
	"preview/stack"
	"strconv"
	"testing"
)

func main() {
	fmt.Println("안녕")
	strconv.Itoa(30)
	stack.StackTest()
	t := testing.T{}
	Map.TestCount(&t)
	fmt.Println(t.Failed())

	// Map의 한계 => 같은 키가 여러번 들어갈 수 있는 맵은 기본적으로 제공하지 않는다.
	// 맵을 여러 고루틴에서 동시에 읽기만 하는 것은 괜찮지만 맵의 구조가 변경괼 수 있는 경우에는 문제가 생긴다.
	// 즉, 스레드 안전하지 않다는 것이다. 락을 이용하거나 다른 라이브러리를 이용해야 한다.
	Map.ExampleCount1()
	Map.ExampleCount2()
	Set.ExampleHasDupeRune()
	method.ExampleAddOne()
	IO.ExampleWriteTo()
	IO.ExampleReadFrom()
	_, err := Error.ExampleError()
	fmt.Println(err)
	LiteralFunction.Example_functionLiteral()
	LiteralFunction.Example_functionLiteralVal()
	HighOrderFunction.ExampleReadFrom_Print()
	HighOrderFunction.ExampleReadFrom_append()
}
