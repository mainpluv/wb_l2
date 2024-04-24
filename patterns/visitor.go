package patterns

import "fmt"

// компонент
type Element interface {
	Accept(Visitor)
}

// конкретный компонент
type ConcreteElementA struct{}

func (c *ConcreteElementA) Accept(v Visitor) {
	v.VisitConcreteElementA(c)
}

// еще один конкретный компонент
type ConcreteElementB struct{}

func (c *ConcreteElementB) Accept(v Visitor) {
	v.VisitConcreteElementB(c)
}

// посетитель
type Visitor interface {
	VisitConcreteElementA(*ConcreteElementA)
	VisitConcreteElementB(*ConcreteElementB)
}

// конкретный посетитель
type ConcreteVisitor struct{}

func (c *ConcreteVisitor) VisitConcreteElementA(e *ConcreteElementA) {
	fmt.Println("Посещение ConcreteElementA")
}

func (c *ConcreteVisitor) VisitConcreteElementB(e *ConcreteElementB) {
	fmt.Println("Посещение ConcreteElementB")
}

// функция посещения
func VisitElements(elements []Element, visitor Visitor) {
	for _, e := range elements {
		e.Accept(visitor)
	}
}

func main() {
	elements := []Element{&ConcreteElementA{}, &ConcreteElementB{}}
	visitor := &ConcreteVisitor{}

	VisitElements(elements, visitor)
}
