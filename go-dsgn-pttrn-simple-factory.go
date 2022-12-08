package main

/* simple factory pattern */

import "fmt"

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
// START : Custom Concrete Product Types

// (abstract) custom product definition
type customProductA struct {
	concreteProduct
}

// free-standing constructor function where variation is declared
func newCustomProductA() iAbstractProduct {
	var item = new(concreteProduct)
	// mutate the basal type
	item.setBehaviour("Custom Product A")
	item.setFeature(1)
	return item
}

// ---

// (abstract) custom product definition
type customProductB struct {
	concreteProduct
}

// free-standing constructor function where variation is declared
func newCustomProductB() iAbstractProduct {
	var item = new(concreteProduct)
	// mutate the basal type
	item.setBehaviour("Custom Product B")
	item.setFeature(2)
	return item
}

// END : Custom Concrete Product Types
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : A Simple Factory Function

// free-standing product selector
func createProduct(productType string) (iAbstractProduct, error) {
	if productType == "customProductA" {
		return newCustomProductA(), nil
	}
	if productType == "customProductB" {
		return newCustomProductB(), nil
	}
	return nil, fmt.Errorf("Wrong product selection.")
}

// END : A Simple Factory Function
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Client Code

func main() {

	myProductA, _ := createProduct("customProductA")
	printDetails(myProductA)

	myProductB, _ := createProduct("customProductB")
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
