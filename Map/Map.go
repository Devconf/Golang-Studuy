package Map

/*
var m map[int]int -> m의 자료형을 맵으로 선언하였지만 초기화 하지 않아 사용 불가능

※맵 초기화 방법
m:=make(map[keyType]valueType)
m:=[keyType]valueType{}

value,ok:=map[key] -> 맵을 읽을때는 두개의 변수로 받게되며 두번째 변수에는 키가 존재여부를 bool값으로 받을수 있다.
만약 키가 있다면 기존에 있는 값은 변경되고 키가 없다면 새로운 키쌍이 추가된다.
*/
import (
	"fmt"
	"reflect"
	"sort"
	"testing"
)

func count(s string, codeCount map[rune]int) {
	for _, r := range s {
		codeCount[r]++
	}
}

func TestCount(t *testing.T) {
	codeCount := map[rune]int{}
	count("가나다나", codeCount)
	if !reflect.DeepEqual(
		map[rune]int{'가': 1, '나': 2, '다': 1},
		codeCount,
	) {
		t.Error("codeCount mismatch:", codeCount)
	}
}

func ExampleCount1() {
	codeCount := map[rune]int{}
	count("가나다나라", codeCount)
	for _, key := range []rune{'가', '나', '다'} { // 이렇게 지정한 키 이외의 키가 들어오면 놓치게됨. 테스트로 다음과 같은 문제를 잡지 못함
		fmt.Println(string(key), codeCount[key])
	}
}

func ExampleCount2() {
	codeCount := map[rune]int{}
	count("가나다나라", codeCount)
	var keys sort.IntSlice
	for key := range codeCount {
		keys = append(keys, int(key))
	}
	sort.Sort(keys)
	for _, key := range keys {
		fmt.Println(string(key), codeCount[rune(key)])
	}
}
