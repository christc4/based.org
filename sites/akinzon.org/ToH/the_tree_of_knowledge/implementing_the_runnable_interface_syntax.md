# Implementing the Runnable Interface Syntax

Up: [Implementing the Runnable Interface](implementing_the_runnable_interface)
Brother(s):
TARGET DECK

```java
class MyClass implements Runnable {
	//properties
	//methods
	public void run() {
		//do something
	}
	public static void main(String[] args){
		Thread myThread = new Thread(MyClass);
		myThread.start(); //need to instantiate Thread to use start()
	}
}


```



































#### Why:
#### How:









