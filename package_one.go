package trial

type Person {
	firstname string
	lastname string
	age int
}

func (p *Person) AddAge(toAdd int) {
	p.age = p.age + toAdd
}

func (p *Person) ChangeFirstName(p newFirstName) {
	p.firstname = newFirstName
}

func (p *Person) ChangeLastName(p newLastName) {
	p.lastname = newLastName
}
