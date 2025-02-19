package main

import (
	"fmt"
	"reflect"
)

// An struct that will implement this interface will be considered a product of factory.
type Interface interface {
	Execute()
}

// A struct that represents the product of a factory.
type ProductA struct {
}

// The implementation of the Interface interface.
func (p *ProductA) Execute() {
	fmt.Println("Execute A")
}

// A struct that represents the product of a factory.
type ProductB struct {
}

// The implementation of the Interface interface.
func (p *ProductB) Execute() {
	fmt.Println("Execute B")
}

// A struct that represents the product of a factory.
type ProductC struct {
}

// The implementation of the Interface interface.
func (p *ProductC) Execute() {
	fmt.Println("Execute C")
}

// The factory struct that will be used to create the products that implement the
// Interface interface.
type Factory struct {
}

// The factory function will always return Interface type. The client will
// always deal with type Interface and its methods, it shouldn't know about
// the implementation of each product.
func (f Factory) CreateProduct(cmd string) Interface {
	switch cmd {
	case "A":
		return &ProductA{}
	case "B":
		return &ProductB{}
	case "C":
		return &ProductC{}
	default:
		return nil
	}
}

func main() {
	f := Factory{}
	for _, cmd := range []string{"A", "B", "C"} {
		p := f.CreateProduct(cmd)
		fmt.Println("The product type", reflect.TypeOf(p))
		p.Execute()
	}

}
