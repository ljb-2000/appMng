package models

type TestType struct {
	Name string
	Id string
}

func PrintHello() (t TestType) {
	t.Name = "luocheng"
	t.Id = "123"
	return
}
