package main

import "fmt"

type State interface {
	Handle(c *Context)
}

type StateA struct {
}

func (s *StateA) Handle(c *Context) {
	fmt.Println("State A")
	c.state = &StateB{}
}

type StateB struct {
}

func (s *StateB) Handle(c *Context) {
	fmt.Println("State B")
	c.state = &StateA{}

}

type Context struct {
	state State
}

func (c *Context) Handle() {
	c.state.Handle(c)
}

func (c *Context) Switch(state State) {
	c.state = state
}

func main() {
	c := &Context{
		state: &StateA{},
	}
	c.Handle()
	c.Handle()
	c.Handle()
	c.Switch(&StateA{})
	c.Handle()
}
