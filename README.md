# go-dsgn-pttrn-simple-factory

> A Go implementation of the Simple Factory design pattern.

Because Go does not support traditional OOP components, such as classes and inheritance, it is not possible to implement the classic Factory Method. Nevertheless, it is still possible to implement a basic Simple Factory pattern within Go.

The Simple Factory Pattern is a creational design pattern that creates objects by calling a free-standing factory function that constructs a particular variant of the specified type. As such, concrete object creation is implemented by derived types rather than by calling a constructor of the primal type itself, thus providing a mechanism to mutate or extend variations of the interested type.

&nbsp;

<p align="center">
  <img src="./go-dsgn-pttrn-simple-factory.svg">
</p>
