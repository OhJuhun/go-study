package shape

type Circle struct {
	Angle  int
	Radius int
}

func (c Circle) GetArea() int  {
	return c.Radius * 3
}

func (c Circle) GetAngles() int {
	return c.Angle
}
