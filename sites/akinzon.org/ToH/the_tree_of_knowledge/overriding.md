# Overriding

Up: [The Object-Oriented Paradigm](the_object-oriented_paradigm)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1202 Programming::Object-Oriented Programming

Overriding is when the behaviour of an [Inheritance|inherited](inheritance|inherited) method defined in a [Superclasses|superclass](superclasses|superclass) is altered in its [subclasses](subclasses).

```java
public class Animal {
	public void makeNoise() {
		System.out.println("Noise");
	}
}

public class Dog extends Animal {
	@override
	public void makeNoise() { // The method signature must be the same
		System.out.println("Woof");
	}
}
```

[Inheritance Trees](inheritance_trees)

`final` prevents overriding from happening.

How does [overriding](overriding) work #flashcard 
[Subclasses](subclasses) of [superclasses](superclasses) will have a [Methods|method](methods|method) with the same [Method Signatures|method signature](method_signatures|method_signature) but with different/added functionality, overriding the original method.
<!--ID: 1705184914844-->



































#### Why:
#### How:









