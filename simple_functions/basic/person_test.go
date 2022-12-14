package basic_test

import (
	"simple_functions/basic"
	"testing"
)

func TestFindPersonById(t *testing.T) {
	people := basic.People{
		basic.Person{
			Id:   0,
			Name: "person 1",
		},

		basic.Person{
			Id:   1,
			Name: "person 2",
		},
	}
	actual, _ := people.FindPersonById(1)
	expected := basic.Person{
		Id:   1,
		Name: "person 2",
	}
	if actual != expected {
		t.Errorf("Expected to find %v but found %v", expected, actual)
	}
}

func TestFindPersonByIdNotFound(t *testing.T) {
	people := basic.People{
		basic.Person{
			Id:   0,
			Name: "person 1",
		},
		basic.Person{
			Id:   1,
			Name: "person 2",
		},
	}
	p, err := people.FindPersonById(2)
	if err == nil {
		t.Errorf("Expected %v but get %v", basic.NotFoundPerson, p)
	}
}
