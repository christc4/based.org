# [Collections](collections) and [Inheritance](inheritance)

Up: [The Object-Oriented Paradigm](the_object-oriented_paradigm)
Brother(s):
TARGET DECK

Since a reference to a class can also store its [subclasses](subclasses), collections can be used to [Generics|generalise](generics|generalise) subclasses by **substitution**. However, calling a method that only belongs to one subclass will not work, as it does not belong to the [Superclasses|superclass](superclasses|superclass).

```java
humans.add(new Man("Rob"));
for(Human h: humans) {
	// do something
}
```



































#### Why:
#### How:









