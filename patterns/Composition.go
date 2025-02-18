package patterns

import "fmt"

type Animal struct {
	Name string
}

func (a *Animal) Voice() {
	fmt.Println("Some sound")
}

type Dog struct {
	Parent *Animal
}

func (d *Dog) Voice() {
	fmt.Println("Woof!")
}

type Cat struct {
	Parent *Animal
}

func (d *Cat) Voice() {
	fmt.Println("Moew!")
}

func CallComposition() {
	dog := Dog{Parent: &Animal{Name: "Buddy"}}
	cat := Cat{Parent: &Animal{Name: "Donna"}}
	fmt.Printf("Dog is %s, ", dog.Parent.Name)
	dog.Voice()
	fmt.Printf("Cat is %s, ", cat.Parent.Name)
	cat.Voice()
}
