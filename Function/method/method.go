package method

import "fmt"

func AddOne(nums []int) {
	for _, val := range nums { // nums의 값을 val로 가져와서 val를 수정 하였을 경우 원 슬라이스인 nums는 변경되지 않는다.
		val++
	}
}

func AddOneByPointer(nums *[]int) {
	for i := 0; i < len(*nums); i++ {
		(*nums)[i]++
	}
}

func ExampleAddOne() {
	n := []int{1, 2, 3, 4}
	AddOne(n)
	fmt.Println(n)
	AddOneByPointer(&n)
	fmt.Println(n)
}
