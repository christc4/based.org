# Limitations of OOP Without Interfaces

Up: [The Object-Oriented Paradigm](the_object-oriented_paradigm)
Brother(s):
TARGET DECK

Given a well-defined inheritance hierarchy like this:

![Pasted image 20240111194106.png](pasted_image_20240111194106.png)

What if we wanted to implement some methods for pets, such as `play()`, and `nuzzle()`? Where would they go?

### Option 1: In the `Animal` class

If they went there, we run into the issue of method behaviour: What does it mean to play with an `Animal`? Is playing with a `Dog` different from playing with a `Cat`? How would this method look?


### Option 2: As an [Abstract Classes|abstract method](abstract_classes|abstract_method) in the `Animal` class

This would require all classes to provide their own implementation - Would you play with a `Tiger?` (I hope not!)


### Option 3: Place method in only the required classes

This breaks the whole point of [polymorphism](polymorphism), since it is not present in all classes (which is already problematic!)

## SO HOW???

> [!tip]- Answer:
> ![Pasted image 20240111194721.png](pasted_image_20240111194721.png)
> [Interfaces](interfaces).

































#### Why:
#### How:









