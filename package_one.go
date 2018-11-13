package trial

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

func (p *Person) AddAge(toAdd int) {
	p.age = p.age + toAdd
}

func (p *Person) ChangeFirstName(newFirstName string) {
	p.firstname = newFirstName
}

func (p *Person) ChangeLastName(newLastName string) {
	p.lastname = newLastName
}

func (p Person) PrintName() {
	fmt.Println("Hello I am ", p.firstname, p.lastname)
}
