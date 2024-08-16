# Methods with parameters but no return type

Up: [Methods](methods)
Brother(s):
TARGET DECK

Parametrised methods are good because they can alter their behaviour based on what parameter has been passed to them.

```java
public class Account{
	int balance = 100;
	public static void main(String[] args){
		Account myAccountObject = new Account();
		myAccountObject.withdrawFiver();
		myAccountObject.withdrawTenner(); 
	}
	public void withdrawFiver(){
		balance = balance - 5; 
	}
	public void withdrawTenner(){ 
		int tenner = 10; 
		balance = balance – 10; 
	} 
}
```

In the code example above, it feels wasteful to have both `withdrawFiver()` and `withdrawTenner()`. Surely there is a method that can emulate them both?

```java
public class Account{
	int balance = 100;
	public static void main(String[] args){
		Account myAccountObject = new Account();
		myAccountObject.withdrawAmount(5);
		myAccountObject.withdrawAmount(10); 
	}
	public void withdrawAmount(int amount){
		balance = balance - amount; 
	}
}
```

There! Much better! The `withdrawAmount(int Amount)` method can do so much more than just withdrawing fivers and tenners, so long as they have an `int` value passed to it, it can get to withdrawing!

> [!info] Methods can have more than one parameter, and they need not be the same type. 

> [!info] Methods must have [variables](variables) of the correct type passed to them or it can lead to interesting side effects.

```java
public class Account{
	int balance = 100;
	public static void main(String[] args){
		Account myAccountObject = new Account();
		myAccountObject.withdraw(“ten pounds”);
		myAccountObject.withdrawAmount('5');
	}
	public void withdrawAmount(int amount){
		balance = balance - amount; 
	}
}

```

passing `"ten pounds"` to `withdrawAmount()` will cause a [Compilers|compiler](compilers|compiler) error since `withdrawAmount()` wants an `int`, not a `String`.

passing `'5'` to `withdrawAmount()` will cause the `char` `'5'` to be converted to its ASCII value, which isn't 5.
































#### Why:
#### How:









