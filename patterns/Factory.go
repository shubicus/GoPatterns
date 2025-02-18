package patterns

import "fmt"

type Car interface {
	Start() string
}

type BMW struct{}
type ZAZ struct{}

func (bmw BMW) Start() string { return "Here is BMW" }

func (zaz ZAZ) Start() string { return "Here is ZAZ" }

func CarFactory(kind string) Car {
	switch kind {
	case "BMW":
		return BMW{}
	case "ZAZ":
		return ZAZ{}
	default:
		return nil
	}
}

func CallFactory(kind string) {
	car := CarFactory(kind)
	if car == nil {
		panic("Car is nil")
	}
	fmt.Println(car.Start())
}
