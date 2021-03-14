package main

import (
	"company"
	"fmt"
	"shape"
	"testlib" //init만을 사용하고 싶은 경우 import _ "testlib"
	"time"
)

var b int = 5

type person struct { //Go는 Class, Object, Inheritance가 없 -> struct로 대신
	name string
	age  int
}

func interfaceAndStructExample() {

	p := person{"jonny", 12}
	println(p.age)
	println(p.name)
	comp := company.Company{Address: "pangyo"}
	comp2 := new(company.Company)
	comp2.Address = "jeju"
	println(&p)
	println(&comp)
	println(comp.Address)
	println(comp2)
	println(comp2.Address)

	comp3 := company.NewCompany()
	comp3.Address = "siheung"
	fmt.Println(comp3)
	comp3.Employee[`jonny`] = `manager`
	fmt.Println(comp3.CalculateAllSalary())
	fmt.Println(comp3.Employee)
	comp3.UpgradeLevel("jonny")
	fmt.Println(comp3.Employee)

	circle := shape.Circle{Radius: 3}
	rectangle := shape.Rectangle{Angle: 4, Length: 5}
	getAnglesAndArea(circle, rectangle)
}
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

func goChannelExample() {
	goChannel := make(chan int)
	// 이를 통해 goroutine에서 데이터를 주고 받는다.
	// 상대편이 준비될 때까지 채널에서 대기함에 따라 별도의 lock을 걸지 않고 데이타를 동기화하는데 사용된다.
	go func() {
		goChannel <- 123
	}()

	var number int
	number = <-goChannel
	println(number)
}

func goChannelExample2() int {
	//Unbuffered Channel
	//하나의 수신자가 데이타를 받을 때까지 송신자가 데이타를 보내는 채널에 묶여 있게 된다.
	done := make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		done <- true
	}()

	// 위의 Go루틴이 끝날 때까지 대기
	<-done

	return 5
}

func bufferedChannelExample() {
	//수신자가 받을 준비가 되어 있지 않을 지라도 지정된 버퍼만큼 데이타를 보내고 계속 다른 일을 수행할 수 있다.
	channel := make(chan string, 1) // buffered channel 생성(type, NumberOfBuffers)
	channel <- "String Channel"
	fmt.Println(<-channel)
	// Unbuffered channel이었다면 별도 수신 goroutine이 없기 때문에 deadlock
}

//수신 전용 channel
func receive(ch <-chan string) {
	data, success := <-ch // 데이터, 수신이 제대로 되었는지
	fmt.Println(data, success)
	if _, success := <-ch; !success {
		println("더이상 데이타 없음.")
	}
}

//송신 전용 channel
func send(ch chan<- string) {
	ch <- "Sending Channel"
}

func channelParameterExample() {
	ch := make(chan string, 1)
	send(ch)
	close(ch)
	//send(ch) //Error
	receive(ch) //Sending Channel, true
	receive(ch) // false
	//defer close(ch) //채널이 닫힌 후에도 수신은 가능하지만, 송신은 불가능
	//for i := range ch { 수신한 만큰 println
	//	println(i)
	//}
}

func goSelectExample() int{
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- true
	}()
	go func() {
		time.Sleep(3 * time.Second)
		ch2 <- true
	}()
	for {
		println("NOT")
		select { //select : 복수의 채널을 동시에 기다린다.
		case <-ch1:
			println("Ch1")
			return 2
		case <-ch2:
			println("Ch2")
			return 1
		}
	}
}
func mapExample() {
	var idMap map[int]string     // = 으로 할당이 불가능(Nil Map)
	map2 := make(map[int]string) // make로 초기화 후에 = 으로 할당이 가능
	println(idMap, map2)
	map2[3] = "heee"
	map2[4] = "hh"
	println(idMap, map2)
	println(map2[3])
	fmt.Println(map2)
	val, exists := map2[3]
	println(val, exists)
	delete(map2, 3)
	println(map2[3])
	val, exists = map2[3]
	println(val, exists)
	for key, value := range map2 {
		println(key, value)
	}
}

func packageExample() {

	song := testlib.GetMusic("Alicia Keys")
	fmt.Println(song)
}

func getAnglesAndArea(shapes ...shape.Shape) {
	for _, s := range shapes {
		fmt.Println(s.GetAngles(), s.GetArea())
	}
}

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	defer println("Good Bye") //해당 함수의 마지막에 실행된다.(에러가 발생하더라도)
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

	// panic("Error received") //에러 Trace
	println("Finished")

	// Go Routine
	// Thread보다 가볍게 concurrency 처리
	// 보통 OS Thread 하나로 처리하고 Go runtime 상에서 관린된다.
	// go channel과  함께 사용하여 통신이 쉽게할 수 있
	go say("Async1")
	go say("Async2")
	go say("Async3")
	//runtime.GOMAXPROCS(NumberOfCpus) //여러 개의 CPU를 가진 머신에서 parallel하게 할 수 있음(Logical CPUs)
	// parallel과 concurrency는 다르다.
	// parallel: 계산의 동시 실행
	// concurrency: 독립적으로 실행되는 process 구성
	// Concurrency 동시에 많은 것들을 처리하는 것이다. Parallelism은 많은 것을 동시에 하는 것이다.
	//time.Sleep(time.Second * 3)

	goChannelExample()

	println("Answer=", goChannelExample2())
	bufferedChannelExample()
	channelParameterExample()
	goSelectExample()
}
