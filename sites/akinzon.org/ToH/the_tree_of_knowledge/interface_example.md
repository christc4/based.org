# Interface Example

Up: [Interfaces](interfaces)
Brother(s):
TARGET DECK

Given the problem in [Limitations of OOP Without Interfaces](limitations_of_oop_without_interfaces), here's how its solved with a `Pet` interface:

```java
public interface Pet {
	void play();
}

public class Dog extends Canine implements Pet {
	public void play() {
		System.out.println(“Dog plays with a ball”);
	}
}

public class Wolf extends Canine {
	
}


```

Since wolves aren't `Pet`s, they don't implement its interface but `Dog`s can, and all of this works!




































#### Why:
#### How:









