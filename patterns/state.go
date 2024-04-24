package patterns

import "fmt"

// интерфейс для состояния
type State interface {
	Handle()
}

// конкретное состояние
type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle() {
	fmt.Println("Обработка A")
}

// конкретное состояние
type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle() {
	fmt.Println("Обработка B")
}

// контекст
type Context struct {
	state State
}

func (c *Context) SetState(state State) {
	c.state = state
}

func (c *Context) Request() {
	if c.state == nil {
		return
	}
	c.state.Handle()
}

func main() {
	context := &Context{}

	stateA := &ConcreteStateA{}
	context.SetState(stateA)
	context.Request()

	stateB := &ConcreteStateB{}
	context.SetState(stateB)
	context.Request()
}
