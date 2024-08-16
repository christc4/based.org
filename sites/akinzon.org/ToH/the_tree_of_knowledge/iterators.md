# Iterators

Up: [Programming Principles](programming_principles), [Design Pattern Examples](design_pattern_examples)
Brother(s):
TARGET DECK

**Back to the [Limitations of ArrayLists](limitations_of_arraylists) example, but this time with iterators:**

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

[Benefits of Iterators](benefits_of_iterators)

[Krazy Iterator casting example](krazy_iterator_casting_example)
































#### Why:
#### How:









