# The Synchronized Keyword

Up: [Using Synchronisation in Java](using_synchronisation_in_java)
Brother(s): [The Synchronized Block](the_synchronized_block)
TARGET DECK


The `synchronized` keyword is used in the [Method Signatures|method signature](method_signatures|method_signature) to declare it as being an entire `synchronized` block in practice.

It is done like this:

```java
public synchronized void someMethod() {
	//some thread-safe code
}
```

For this way of adding synchronisation, the [Objects|object](objects|object) this method is invoked on is the [Thread Locks|lock](thread_locks|lock).

































#### Why:
#### How:









