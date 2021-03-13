package main

import (
	"fmt"
	"testlib" //init만을 사용하고 싶은 경우 import _ "testlib"
)

var b int = 5

func switchCaseExample(i int) {
	switch i {
	case 4:
		println("4 이상")
		fallthrough
	case 3:
		println("3 이상")
		fallthrough // fallthrough를 사용했으므로 아래 case를 모두 실행
	case 2:
		println("2 이상")
		fallthrough // 기본적으로 case가 맞으면 break가 됨. 그걸 무시하는 keyword
	case 1, 0:
		println("1 이상") // 마지막 case에는 fallthrough를 사용할 수 없음
	}

	switch {
	case i >= 5:
		println("i is bigger than 5")
	case i < 5:
		println("i is less than 5")
	}
}

func goOnlyHasFor(i int) {
	for j := 0; j < i; j++ {
		println(j)
	}

	for i >= 0 {
		i--
	}
	println(i)
	names := []string{"홍길동", "이순신", "강감찬"}

	for name, ff := range names {
		println(name, ff)
	}
}

func passByValue(msg string) {
	println("in passByValue", msg, &msg)
	msg = "01098765432"
	println("value updated", msg, &msg)
}

func passByReference(msg *string) {
	println("in passByReference", msg, &msg, *msg)
	*msg = "jonny"
	println("value updated", msg, &msg, *msg)
}

func variadicFunction(msg ...string) {
	println(msg)

	for _, m := range msg {
		print(m, " ")
	}
	println()
}

func returnMultiValues() (int, string, bool) {
	return 1, "String", false
}

func namedReturnMultiValues() (d int, e string, f bool) {
	d = 2
	e = "STRING"
	f = true
	return
}

func allocateExample() {
	var number int
	var long int64 = 9223372036854775807
	var unsignedInt uint32 = 4
	var float float64 = 2.
	var double complex128 = 34.555
	var str string = "HELLO" // immutable
	var longStr string = `hello,\n //ddd
world`
	var noneType = 4
	const constNum = 33
	var f bool
	t := true
	var b byte = 5
	var r rune = 55
	println(number, long, unsignedInt, float, double, str, longStr, noneType, constNum, f, t, b, r)

	st := "ABCDEFG"
	bytes := []byte(st)
	str2 := string(bytes)
	println(bytes, str2)
}

func pointerExample() {
	var k int = 10
	var p = &k
	println("Pointer Example:", p, *p)
}

func ifExample(boolean bool) {
	if boolean { //block 필수
		println("This is True")
	} else { //else if, else는 무조건 같은 라인
		println("This is False")
	}
}

// calculator 원형 정
type calculator func(int, int) int

// calculator 원형 사용
func calc(f calculator, a int, b int) int {
	result := f(a, b)
	return result
}

// Closure 함수
func closureExample() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func arrayExample() {
	var arr = [...]int{5, 6, 7}
	for _, num := range arr {
		println(num)
	}
}

func sliceExample() {
	s := make([]int, 5, 10) //Nil Slice
	println(len(s), cap(s))
	for _, num := range s {
		println(num)
	}

	arr := []int{0, 1}

	println(arr)
	// 하나 확장
	arr = append(arr, 2) // 0, 1, 2
	println(arr)
	// 복수 요소들 확장
	arr = append(arr, 3, 4, 5) // 0,1,2,3,4,5
	println(arr)
	fmt.Println(arr)

	ss := []int{1, 2, 3, 4, 5}
	println(ss)
	ss = ss[2:3]
	println(ss)
}

func mapExample() {
	var idMap map[int]string // = 으로 할당이 불가능(Nil Map)
	map2 := make(map[int]string) // make로 초기화 후에 = 으로 할당이 가능
	println(idMap, map2)
	map2[3]="heee"
	map2[4]="hh"
	println(idMap,map2)
	println(map2[3])
	fmt.Println(map2)
	val, exists := map2[3]
	println(val,exists)
	delete(map2,3)
	println(map2[3])
	val, exists = map2[3]
	println(val,exists)
	for key,value := range map2{
		println(key,value)
	}
}

func main() {
	println("Hello, World")
	allocateExample()
	pointerExample()
	ifExample(true)
	switchCaseExample(3)
	goOnlyHasFor(4)
	name := "Juhun"
	println(name, &name)
	passByReference(&name)
	println(name, &name)

	mobile := "01012345678"
	println(mobile, &mobile)
	passByValue(mobile)
	println(mobile, &mobile)

	variadicFunction("My")
	variadicFunction("My", "Name")
	variadicFunction("My", "Name", "Is")

	var a, b, c = returnMultiValues()
	println(a, b, c)

	var f, d, e = namedReturnMultiValues()
	println(f, d, e)

	// Anonymous Function
	// 변수명이 함수명과 동일하게 취급
	// go에서 함수는 first-class function
	sum := func(n ...int) int { //익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	println(sum(1, 2, 3, 4, 5))
	println(calc(func(a int, b int) int {
		return a + b
	}, 1, 2))

	firstClosure := closureExample()
	println(firstClosure())
	println(firstClosure())

	secondClosure := closureExample()
	println(secondClosure())
	println(secondClosure())
	arrayExample()
	sliceExample()
	mapExample()
	song:= testlib.GetMusic("Alicia Keys")
	fmt.Println(song)
}
