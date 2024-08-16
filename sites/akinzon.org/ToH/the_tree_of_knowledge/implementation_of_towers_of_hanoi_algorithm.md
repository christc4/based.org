# Implementation of Towers of Hanoi Algorithm

Up: [Towers of Hanoi](towers_of_hanoi)
Brother(s):
TARGET DECK

```java
hanoi(n, A, B, C)
{
	if (n > 0) {
	hanoi(n-1, A, C, B);
	move(A, C);
	hanoi(n-1, B, A, C);
	}
}
```

This is good we only need to move the largest disk from peg A to peg C once, moving all other disks to peg B first. If this works for the $n-1$ problem then it works for $n$, 



































#### Why:
#### How:









