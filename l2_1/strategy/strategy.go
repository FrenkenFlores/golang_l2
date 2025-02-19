package main

import "fmt"

type Interface interface {
	Execute()
}

type CommandA struct {
}

func (ca *CommandA) Execute() {
	fmt.Println("Execute command A")
}

type CommandB struct {
}

func (cb *CommandB) Execute() {
	fmt.Println("Execute command B")
}

type CommandC struct {
}

func (cc *CommandC) Execute() {
	fmt.Println("Execute command C")
}

type Object struct {
}

func (ob *Object) Execute(cmd Interface) {
	cmd.Execute()
}

func main() {
	ob := Object{}
	ob.Execute(&CommandA{})
	ob.Execute(&CommandB{})
	ob.Execute(&CommandC{})
}
