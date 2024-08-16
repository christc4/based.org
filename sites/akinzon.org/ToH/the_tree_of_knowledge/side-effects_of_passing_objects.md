# Side-Effects of Passing Objects

Up: [Objects](objects)
Brother(s):
TARGET DECK

As mentioned in the legendary battle of [Primitives vs Objects](primitives_vs_objects), objects are never actually passed, but a reference to that object in [RAM|main memory](ram|main_memory) is instead.

So how can this go wrong?

Look at the following code:

```java
public void updateAge(Person person) {
	System.out.println(person.getAge()); // person is 56
	person.incrementAge(1); // now they are 57
	Person personB = person; // person's reference is copied to personB
	personB.incrementAge(2) // person is now 59!!!
}
```

| Stack | Heap |
| ---- | ---- |
| person | 0xff343 |
| personB |  |
Both `person` and `personB` have a reference to memory location `0xff343`, so updating one updates the other.



































#### Why:
#### How:









