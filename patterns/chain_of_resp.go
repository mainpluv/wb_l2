package patterns

import "fmt"

// интерфейс хендлера
type Handler interface {
	HandleRequest(request int)
	SetNext(handler Handler)
}

// конкретный хендлер 1
type ConcreteHandler1 struct {
	next Handler
}

func (c *ConcreteHandler1) HandleRequest(request int) {
	if request < 10 {
		fmt.Println("Запрос обрабатывается с помощью concretehandler1")
	} else if c.next != nil {
		c.next.HandleRequest(request)
	}
}

func (c *ConcreteHandler1) SetNext(handler Handler) {
	c.next = handler
}

// конкретный хендлер 2
type ConcreteHandler2 struct {
	next Handler
}

func (c *ConcreteHandler2) HandleRequest(request int) {
	if request < 20 {
		fmt.Println("Запрос обрабатывается ConcreteHandler2")
	} else if c.next != nil {
		c.next.HandleRequest(request)
	}
}

func (c *ConcreteHandler2) SetNext(handler Handler) {
	c.next = handler
}

// клиент
func main() {
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}

	handler1.SetNext(handler2)

	requests := []int{5, 15, 25}

	for _, request := range requests {
		handler1.HandleRequest(request)
	}
}
