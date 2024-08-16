# Methods with no parameters and no return type

Up: [Methods](methods)
Brother(s):
TARGET DECK

These type of methods simply do a job.

**Examples:**

```java
public void sayHello() {
	System.out.println("Hallo");
}

public void onePlusOne() {
	int a = 1;
	System.out.println(a + 1);
}
```

**How to call such a method from the same [Classes|class](classes|class) and from a different class:**

```java
public void onePlusOne() {
	sayHello(); // same class
	person.printName(); // object of a different class
	int a = 1;
	System.out.println(a + 1); 
}
```































#### Why:
#### How:









