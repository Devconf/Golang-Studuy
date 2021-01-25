package Set

import "fmt"

// 상수 시간에 키의 존재를 확인 할 수 있는 집압은 따로 제공하지 않음
// 코드도 깔끔하고 좋지만 불필요한 bool값 때문에 메모리를 차지한다.
func hasDupeRune1(s string) bool {
	runeSet := map[rune]bool{}
	for _, r := range s {
		if runeSet[r] {
			return true
		}
		runeSet[r] = true
	}
	return false
}

// 메모리 차지를 해결하기위해 빈구조체를 값으로 사용하면 된다.
// 만약 구조체가 존재하면 r이 true 가 나오고 없으면 false가 나온다.
func hasDupeRune2(s string) bool {
	runeSet := map[rune]struct{}{}
	for _, r := range s {
		if _, exists := runeSet[r]; exists {
			return true
		}
		runeSet[r] = struct{}{} // struct{}는 자료형이다.
	}
	return false
}

func ExampleHasDupeRune() {
	fmt.Println(hasDupeRune1("숨바꼭질"))
	fmt.Println(hasDupeRune2("다시합시다"))
}
