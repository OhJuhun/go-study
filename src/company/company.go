package company

type Company struct { //대문자로 시작해야 외부 package에서 사용 가
	name     string
	Address  string //field도 대문자여야만 외부에서 사용이 가능
	Employee map[string]string
}

func NewCompany() *Company { // 요걸 생성자로 생각하면 된다.
	println("Company Created")
	n := "J"
	e := map[string]string{}
	comp := Company{name: n, Employee: e}
	return &comp
}

// method는 아래처럼 분리되어 정의
func (c Company) CalculateAllSalary() int { //Value Receiver
	return 100 * len(c.Employee)
}

func (c *Company) UpgradeLevel(name string) { //Pointer Receiver
	c.Employee[name] = "CLEVEL"
}
