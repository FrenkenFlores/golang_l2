package main

import "fmt"

// Unit interface
type Unit interface {
	accept(v *Visitor)
}

// Unit element
type UnitA struct {
	a int
}

// Interface implementation
func (ua *UnitA) accept(v Visitor) {
	v.VisitA(ua)
}

// Unit visitor interface
type Visitor interface {
	VisitA(ua *UnitA)
}

// Unit visitor
type VisitorA struct {
}

// Unit visitor interface implementation
func (v *VisitorA) VisitA(ua *UnitA) {
	ua.a = 1
}

func main() {
	// Create unit element
	ua := &UnitA{a: 0}
	// Print init value
	fmt.Println(ua.a)
	// Create visitor element
	va := &VisitorA{}
	// Accept visitor
	ua.accept(va)
	// Print Changed value
	fmt.Println(ua.a)
}
