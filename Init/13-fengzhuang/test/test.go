package test

type Human struct {
	Sex  int
	name string
}

func (h *Human) Walk() {
	println("human", h.name, "is walking...")
}
func (h *Human) Eat() {
	println("human", h.name, " is eating...")
}

type Superman struct {
	// 继承
	Human
	Level int
}

func (s *Superman) Eat() {
	println("superman eat...")
}

func (s *Superman) Fly() {
	println("superman", s.name, " is flying...")
}
