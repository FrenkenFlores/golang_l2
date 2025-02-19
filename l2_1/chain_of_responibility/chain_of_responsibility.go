package main

import "fmt"

type Handler interface {
	Handle(request string) string
	SetNextHandler(h Handler) Handler
}

type BaseHandler struct {
	next Handler
}

func (h *BaseHandler) Handle(request string) string {
	if h.next != nil {
		return h.next.Handle(request)
	}
	return ""
}

func (h *BaseHandler) SetNextHandler(nh Handler) Handler {
	h.next = nh
	return nh
}

type ConcreteHandlerA struct {
	BaseHandler
}

func (cha *ConcreteHandlerA) Handle(request string) string {
	if request == "A" {
		return "Handeling request A"
	}
	return cha.next.Handle(request)
}

type ConcreteHandlerB struct {
	BaseHandler
}

func (cha *ConcreteHandlerB) Handle(request string) string {
	if request == "B" {
		return "Handeling request B"
	}
	return cha.next.Handle(request)
}

type ConcreteHandlerC struct {
	BaseHandler
}

func (cha *ConcreteHandlerC) Handle(request string) string {
	if request == "C" {
		return "Handeling request C"
	}
	return cha.next.Handle(request)
}

func main() {
	ha := &ConcreteHandlerA{}
	hb := &ConcreteHandlerB{}
	hc := &ConcreteHandlerC{}
	ha.SetNextHandler(hb).SetNextHandler(hc)
	for _, i := range []string{"A", "B", "C"} {
		response := ha.Handle(i)
		fmt.Println(response)
	}

}
