package basic

import "errors"

type Person struct {
	Id   int
	Name string
}

type People []Person

var NotFoundPerson = errors.New("can not find person")

func (p People) FindPersonById(id int) (Person, error) {
	for _, person := range p {
		if person.Id == id {
			return person, nil
		}
	}
	return Person{}, NotFoundPerson
}
