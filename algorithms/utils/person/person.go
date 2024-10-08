package person

import . "errors"

type Person struct {
	age int
}

func NewPerson(age int) (*Person, error) {
	if age < 1 {
		return nil, New("A person is at least 1 years old")
	}

	return &Person{
		age: age,
	}, nil
}

func (p *Person) Older(other *Person) bool {
	return p.age > other.age
}

func (p *Person) Age() int {
	return p.age
}
