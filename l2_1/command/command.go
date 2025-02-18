package main

import "fmt"

type Command interface {
	execute()
}

type CommandA struct {
	m *Main
	a string
}

func (c *CommandA) execute() {
	c.m.a = c.a
	fmt.Println("Command", c.m.a)
}

type CommandB struct {
	m *Main
	b string
}

func (c *CommandB) execute() {
	c.m.b = c.b
	fmt.Println("Command", c.m.b)
}

type CommandC struct {
	m *Main
	c string
}

func (c *CommandC) execute() {
	c.m.c = c.c
	fmt.Println("Command", c.m.c)
}

type Main struct {
	a string
	b string
	c string
}

func (m *Main) CommandA(a string) Command {
	return &CommandA{
		a: a,
		m: m,
	}
}

func (m *Main) CommandB(b string) Command {
	return &CommandB{
		b: b,
		m: m,
	}
}

func (m *Main) CommandC(c string) Command {
	return &CommandC{
		c: c,
		m: m,
	}
}

func main() {
	m := Main{}
	cmds := []Command{}
	cmds = append(cmds,
		m.CommandA("A"),
		m.CommandB("B"),
		m.CommandC("C"),
	)
	for _, cmd := range cmds {
		cmd.execute()
	}
}
