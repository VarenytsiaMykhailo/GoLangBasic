package person

import (
	"fmt"
)

func NewPerson(id int, name, secret string) *Person {
	return &Person{
		ID:     id,
		Name:   name,
		secret: secret,
	}
}

func GetSecret(p *Person) string {
	return p.secret
}

func printSecret(p *Person) {
	fmt.Println(p.secret)
}
