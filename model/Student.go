package model

type Student struct {
	Name    string
	Address string
}

func New(name string, address string) *Student {
	return &Student{
		Name:    name,
		Address: address,
	}
}
