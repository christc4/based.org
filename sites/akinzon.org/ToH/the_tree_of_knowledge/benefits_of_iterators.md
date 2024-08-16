# Benefits of Iterators

Up: [Iterators](iterators)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1202 Programming::Programming Principles

Iterators make code look neater, and neat code is easier to understand.

Iterators also decouple the loop from the collection (As you can see, the iterator doesn't even reference the original `ArrayList`).

As a result, we can pass the iterator to a [methods|method](methods|method), which doesn't even need to know what the [Collections|collection](collections|collection) is! 

This is really useful if the underlying type can be an [ArrayLists|ArrayList](arraylists|arraylist) or [Arrays|Array](arrays|array), which have different ways of determining size (`.size()` and `.length` respectively) and different ways of indexing.

```java
ArrayList kennel = new ArrayList<Dog>;

kennel.add(new Dog("Bingo"));
kennel.add(new Dog("Bango"));
kennel.add(new Dog("Bongo"));

Iterator it = kennel.iterator();
while(it.hasNext()) { // checks if there's any more items left 
	it.next().bark(); // next goes to the next item
}
```

What are the [benefits of Iterators](benefits_of_iterators)? #flashcard 
They decouple the loop from the collection, allowing you to use the same loop on different types of collection.
It protects encapsulated information in the collection object.
<!--ID: 1705184914938-->





































#### Why:
#### How:









