# Dynamic Binding

Up: [The Object-Oriented Paradigm](the_object-oriented_paradigm)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1202 Programming::Object-Oriented Programming

From [Collections and Inheritance](collections_and_inheritance) we know that substitution is used so that a generic collection can hold [objects](objects) that are [subclasses](subclasses) of the declared collection's type.

This can be taken a step further with **Dynamic Binding**, where we can call a method defined in the generic class, but it will use the behaviour of the specific object's version of that method.

```java
ArrayList animals;
animals = new ArrayList();
animals.add(new Animal()); 
animals.add(new Dog()); 
animals.add(new Cat()); 
animals.add(new Elephant()); 
for(Animal a : animals) {
	a.sleep(); 
	a.eat(); 
	a.roam();
}
```

`sleep()`, `eat()` and `roam()` are all defined in the `Animal` superclass (or else there would be a compiler error), but the behaviour used any time they are called is based on which object `(Dog, Cat, Elephant)` has its method called.

So even though we are iterating through an `Animal` collection, we are using the behaviour of its subclasses where possible - `Dog`s, `Cat`s and `Elephant`s!!

How does [dynamic binding](dynamic_binding) work #flashcard 
Dynamic binding is when [Methods](methods) of a generic [Classes|class](classes|class) are called, but the actual behaviour of that method is defined by the specific [Objects|object's](objects|object's) version of that method.
<!--ID: 1705184914910-->




































#### Why:
#### How:









