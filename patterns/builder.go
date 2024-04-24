package patterns

import "fmt"

// продукт
type Pizza struct {
	Dough   string
	Sauce   string
	Topping string
}

// строитель
type PizzaBuilder interface {
	SetDough() PizzaBuilder
	SetSauce() PizzaBuilder
	SetTopping() PizzaBuilder
	Build() Pizza
}

// конкретный строитель
type HamPizzaBuilder struct {
	pizza Pizza
}

func (b *HamPizzaBuilder) SetDough() PizzaBuilder {
	b.pizza.Dough = "тесто"
	return b
}

func (b *HamPizzaBuilder) SetSauce() PizzaBuilder {
	b.pizza.Sauce = "томатный соус"
	return b
}

func (b *HamPizzaBuilder) SetTopping() PizzaBuilder {
	b.pizza.Topping = "ветчина и сыр"
	return b
}

func (b *HamPizzaBuilder) Build() Pizza {
	return b.pizza
}

// директор
type Cook struct {
	builder PizzaBuilder
}

func NewCook(builder PizzaBuilder) *Cook {
	return &Cook{builder: builder}
}

func (c *Cook) MakePizza() Pizza {
	return c.builder.SetDough().SetSauce().SetTopping().Build()
}

func main() {
	builder := &HamPizzaBuilder{}
	c := NewCook(builder)
	pizza := c.MakePizza()
	fmt.Println(pizza)
}
