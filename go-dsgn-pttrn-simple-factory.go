package main

/* simple factory pattern */

import (
	"fmt"
)

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Abstract Product Definition (Interface)

// product behaviours
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
// START : Concrete Product A Definition

// the concrete product type
type concreteProductA struct {
	behaviourA string
	featureA   int
}

// the concrete product implementation:
// ---------------------------------

func (item *concreteProductA) setBehaviour(behaviour string) {
	item.behaviourA = behaviour
}

func (item *concreteProductA) getBehaviour() string {
	return item.behaviourA
}

func (item *concreteProductA) setFeature(feature int) {
	item.featureA = feature
}

func (item *concreteProductA) getFeature() int {
	return item.featureA
}

// END : Concrete Product A Definition
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Concrete Product A Definition

// the concrete product type
type concreteProductB struct {
	behaviourB string
	featureB   int
}

// the concrete product implementation:
// ---------------------------------

func (item *concreteProductB) setBehaviour(behaviour string) {
	item.behaviourB = behaviour
}

func (item *concreteProductB) getBehaviour() string {
	return item.behaviourB
}

func (item *concreteProductB) setFeature(feature int) {
	item.featureB = feature
}

func (item *concreteProductB) getFeature() int {
	return item.featureB
}

// END : Concrete Product A Definition
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Simple Factory Function

// a product selector with some mutation
func createProduct(productType string, behaviour string, feature int) iAbstractProduct {
	switch productType {
	case "concreteProductA":
		return &concreteProductA{
			behaviourA: behaviour,
			featureA:   feature,
		}
	case "concreteProductB":
		return &concreteProductB{
			behaviourB: behaviour,
			featureB:   feature,
		}
	default:
		return nil
	}
}

// END : Simple Factory Function
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

// v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v v
// START : Client Code

func main() {

	myProductA := createProduct("concreteProductA", "myProductA", 123)
	printDetails(myProductA)

	myProductB := createProduct("concreteProductB", "myProductB", 456)
	printDetails(myProductB)

	// can't access attributes directly but can use
	// methods to modify after object creation:
	myProductB.setBehaviour("Need unique change here.")
	myProductB.setFeature(69)
	printDetails(myProductB)
}

// Test Function
func printDetails(product iAbstractProduct) {
	fmt.Printf("Behaviour: %s\n", product.getBehaviour())
	fmt.Printf("Feature  : %d\n", product.getFeature())
	fmt.Printf("\n")
}

// END : Client Code
// ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^ ^
// - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -
