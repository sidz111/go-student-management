package model

type Student struct {
	Id      int
	Name    string
	Address string
}

func New(id int, name string, address string) *Student {
	return &Student{
		Id:      id,
		Name:    name,
		Address: address,
	}
}
