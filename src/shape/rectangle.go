package shape

type Rectangle struct {
	Angle  int
	Length int
}

func (r Rectangle) GetArea() int {
	return r.Length * r.Length
}

func (r Rectangle) GetAngles() int{
	return r.Angle
}