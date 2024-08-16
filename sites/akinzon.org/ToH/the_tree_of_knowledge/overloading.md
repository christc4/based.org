# Overloading

Up: [The Object-Oriented Paradigm](the_object-oriented_paradigm)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester I::1202 Programming::Object-Oriented Programming

To understand overloading, you must first understand [Method Signatures](method_signatures). 

Overloading takes advantage of being able to change a method signature simply by changing the parameters. You can call virtually the same method with different parameters and the method's behaviour will change!

```java
public class CashRegister{
	public static void main(String[] args) {
		CashRegister reg = new CashRegister();
		float p1 = 10.0f;
		p1 = reg.calcVAT(p1); // calls the first calcVAT
		p1 = reg.calcVAT(p1, 0.175f); // calls the second calcVAT
		System.out.println(p1);
	}
	public float calcVAT(float price) {
		return price * 1.2f;
	}
	public float calcVAT(float price, float rate) {
		return price * (1.0f + rate); 
	}
}
```

What has to be changed about a [Methods|method](methods|method) for it to count as [Overloading](overloading), and what doesn't work #flashcard 
For overloading to work, the parameter types or number of parameters have to change.
Changing the return type or the names of parameters doesn't change the [Method Signatures|method signature](method_signatures|method_signature).
<!--ID: 1705184914855-->

































#### Why:
#### How:









