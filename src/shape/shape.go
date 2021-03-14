package shape

// interface는 method만 가질 수 있다.
// 이 Method를 모두 구현한 struct가 이를 상속한 것으로 판단된다.
// 뭘 상속했는지 한 눈에 판단이 어려움..?
type Shape interface {
	GetArea() int
	GetAngles() int
}
// 모든 Type을 나타내기 위해 interface{} (empty Interface) 사