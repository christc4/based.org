# Limitations of [ArrayLists](arraylists)

Up: [Java Principles](java_principles)
Brother(s):
TARGET DECK

While ArrayLists aren't restricted by only being able to store one type, this freedom can result in problems in the program further on.

**Here's an example:**

```java
ArrayList kennel = new ArrayList();

kennel.add(new Dog("Bingo"));
kennel.add(new Dog("Bango"));
kennel.add(new Dog("Bongo"));
kennel.add(new cat("Robbie")); // cats can't bark!!!

for(int i; i < kennel.size(); i++) {
	kennel.get(i).bark();
}
```

The above code will not [Compilers|compile](compilers|compile) because Java does not know what is in the `ArrayList` and if `bark()` can be called on it

`instanceof` can be used to mitigate this problem, but this is not advised, since it is bad programming practice.

So how do we actually solve it?

[Generics](generics)





























#### Why:
#### How:









