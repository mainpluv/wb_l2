package patterns

import "fmt"

// интерфейс для стратегии
type Strategy interface {
	DoOperation(int, int) int
}

// конкретная стратегия
type AddOperation struct{}

func (a *AddOperation) DoOperation(num1, num2 int) int {
	return num1 + num2
}

// конкретная стратегия
type SubtractOperation struct{}

func (s *SubtractOperation) DoOperation(num1, num2 int) int {
	return num1 - num2
}

// контекст
type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	if c.strategy == nil {
		return 0
	}
	return c.strategy.DoOperation(num1, num2)
}

func main() {
	context := &Context{}

	addStrategy := &AddOperation{}
	context.SetStrategy(addStrategy)
	fmt.Println("10 + 5 =", context.ExecuteStrategy(10, 5))

	subtractStrategy := &SubtractOperation{}
	context.SetStrategy(subtractStrategy)
	fmt.Println("10 - 5 =", context.ExecuteStrategy(10, 5))
}
