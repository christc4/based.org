# Methods with parameters and a return type

Up: [Methods](methods)
Brother(s):
TARGET DECK

These methods are highly versatile in their usage, they can be used to calculate new values to be assigned to [variables](variables), update existing variables, and more!

```java
public class AreaCalculator {
	int length;
	int width;
	
	public AreaCalculator(int length, int width) {
		this.length = length;
		this.width = width;
	}
	public int calculateVolume(int height) {
		return length * width * height;
	}
	public static void main(String[] args) {
		AreaCalculator areaCalc = new AreaCalculator(3, 4);
		int volume = areaCalc.calculateVolume(5);
	}
	
}
```

[The main method](the_main_method) can assign the result of `calculateVolume()` to a variable



































#### Why:
#### How:









