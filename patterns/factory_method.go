package patterns

import "fmt"

// интерфейс для продукта
type Product interface {
	Use() string
}

// конкретный продукт
type ConcreteProductA struct{}

func (p *ConcreteProductA) Use() string {
	return "Using product A"
}

// конкретный продукт
type ConcreteProductB struct{}

func (p *ConcreteProductB) Use() string {
	return "Using product B"
}

// интерфейс создателя
type Creator interface {
	CreateProduct() Product
}

// конкретный создатель
type ConcreteCreatorA struct{}

func (c *ConcreteCreatorA) CreateProduct() Product {
	return &ConcreteProductA{}
}

// конкретный создатель
type ConcreteCreatorB struct{}

func (c *ConcreteCreatorB) CreateProduct() Product {
	return &ConcreteProductB{}
}

func main() {
	creatorA := &ConcreteCreatorA{}
	productA := creatorA.CreateProduct()
	fmt.Println(productA.Use())

	creatorB := &ConcreteCreatorB{}
	productB := creatorB.CreateProduct()
	fmt.Println(productB.Use())
}
