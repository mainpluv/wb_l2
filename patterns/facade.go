package patterns

import "fmt"

// какой-то компонент A
type ComponentA struct{}

func (a *ComponentA) MethodA() {
	fmt.Println("Выполнение метода A компонента A")
}

// еще один компонент B
type ComponentB struct{}

func (b *ComponentB) MethodB() {
	fmt.Println("Выполнение метода B компонента B")
}

// упрощенный интерфейс
type Facade struct {
	componentA *ComponentA
	componentB *ComponentB
}

func NewFacade() *Facade {
	return &Facade{
		componentA: &ComponentA{},
		componentB: &ComponentB{},
	}
}

func (f *Facade) DoSomething() {
	fmt.Println("Фасад делает что-то полезное:")
	f.componentA.MethodA()
	f.componentB.MethodB()
}

func main() {
	facade := NewFacade()
	facade.DoSomething()
}
