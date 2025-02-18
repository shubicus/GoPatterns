package main

import "fmt"

type Cake struct {
	Sweets    string
	Chocolate bool
	Weight    int
}

type CakeBuilder interface {
	WithSweets(value string) CakeBuilder
	WithChocolate(value bool) CakeBuilder
	WithWeight(value int) CakeBuilder
	Build() Cake
}

type CakeBuilderImpl struct {
	product Cake
}

func (b *CakeBuilderImpl) WithSweets(value string) CakeBuilder {
	b.product.Sweets = value
	return b
}

func (b *CakeBuilderImpl) WithChocolate(value bool) CakeBuilder {
	b.product.Chocolate = value
	return b
}

func (b *CakeBuilderImpl) WithWeight(value int) CakeBuilder {
	b.product.Weight = value
	return b
}

func (b *CakeBuilderImpl) Build() Cake {
	return b.product
}

func CallBuilder() {
	builder := &CakeBuilderImpl{}
	product := builder.
		WithSweets("Marmelade").
		WithChocolate(true).
		WithWeight(1).
		Build()

	fmt.Printf("Cake: %+v\n", product)
}
