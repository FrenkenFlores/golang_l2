package main

import "fmt"

type Product struct {
	a int
	b bool
	c string
}

type Builder struct {
	p Product
}

func (b *Builder) SetA(aValue int) *Builder {
	b.p.a = aValue
	return b
}

func (b *Builder) SetB(bValue bool) *Builder {
	b.p.b = bValue
	return b
}

func (b *Builder) SetC(cValue string) *Builder {
	b.p.c = cValue
	return b
}

func (b *Builder) Build() Product {
	return b.p
}

func main() {
	builder := Builder{p: Product{}}
	product := builder.SetA(0).SetB(true).SetC("Hello, world!").Build()
	fmt.Println(product.a, product.b, product.c)
}
