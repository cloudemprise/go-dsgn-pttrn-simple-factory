package main

/* simple factory pattern */

import (
	"fmt"
)

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Abstract Product Definition (Interface)

// require product behaviours
type iAbstractProduct interface {
	setBehaviour(behaviour string)
	getBehaviour() string
	setFeature(feature int)
	getFeature() int
}

// END : Abstract Product Definition (Interface)
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Product Definition

// the concrete product type
type concreteProduct struct {
	behaviour string
	feature   int
}

// the concrete product implementation:
// ---------------------------------

func (item *concreteProduct) setBehaviour(behaviour string) {
	item.behaviour = behaviour
}

func (item *concreteProduct) getBehaviour() string {
	return item.behaviour
}

func (item *concreteProduct) setFeature(feature int) {
	item.feature = feature
}

func (item *concreteProduct) getFeature() int {
	return item.feature
}

// END : Concrete Product Definition
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Custom Concrete Product Factory Functions

// (abstract) custom product definition
type customProductA struct {
	concreteProduct
}

// factory function : one way of doing it
func newCustomProductA() iAbstractProduct {
	return &customProductA{
		concreteProduct{
			behaviour: "Special default",
			feature:   0,
		},
	}
}

// ---

// (abstract) custom product definition
type customProductB struct {
	concreteProduct
}

// factory function : a different way of doing it
func newCustomProductB() iAbstractProduct {
	// var item = &customProductB{}
	var item = new(customProductB)
	item.behaviour = "Different default value"
	item.feature = 0
	return item
}

/* A 3rd way of doing it
// free-standing factory function where variation is declared
func newCustomProductB() iAbstractProduct {
	var item = new(concreteProduct)
	// mutate the basal type
	item.setBehaviour("Custom Product B")
	item.setFeature(2)
	return item
} */

// END : Custom Concrete Product Factory Functions
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Product Selector

// a product selector with some mutation
func createProduct(productType string) iAbstractProduct {
	switch productType {
	case "customProductA":
		product := newCustomProductA()
		// mutate the default values
		product.setBehaviour("Custom Product A")
		product.setFeature(123)
		return product
	case "customProductB":
		product := newCustomProductB()
		// mutate the default values
		product.setBehaviour("Custom Product B")
		product.setFeature(456)
		return product
	default:
		return nil
	}

}

// END : Product Selector
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Client Code

func main() {

	myProductA := createProduct("customProductA")
	printDetails(myProductA)

	myProductB := createProduct("customProductB")
	myProductB.setBehaviour("Need unique change here.")
	myProductB.setFeature(69)
	printDetails(myProductB)
}

// Test Function
func printDetails(product iAbstractProduct) {
	fmt.Printf("Behaviour: %s\n", product.getBehaviour())
	fmt.Printf("Feature  : %d\n", product.getFeature())
}

// END : Client Code
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
