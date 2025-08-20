package taskdua

import "fmt"

type Person struct {
	Name     string
	Address  string
	PhoneNum string
}

func NewPerson(name, address, phoneNum string) *Person {
	return &Person{
		Name: name, Address: address, PhoneNum: phoneNum,
	}
}

func (p *Person) GetPerson() {
	fmt.Printf("name: %s\naddress: %s\nphone: %s\n", p.Name, p.Address, p.PhoneNum)
}

func (p *Person) Greet() {
	fmt.Printf("Hello, %s\n", p.Name)
}

func (p *Person) SetNewName(newName string) {
	p.Name = newName
}
