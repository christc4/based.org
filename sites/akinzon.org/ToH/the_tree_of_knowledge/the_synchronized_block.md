# The Synchronized Block

Up: [Using Synchronisation in Java](using_synchronisation_in_java)
Brother(s): [The Synchronized Keyword](the_synchronized_keyword)
TARGET DECK

To use the `synchronized` block, you wrap a block of code in it like this: 

```java
...
synchronized(obj) {
	//some thread-safe code
}
...
```

The `obj` is the [Objects|object](objects|object) used as a [Thread Locks|lock](thread_locks|lock). It's more like a key functionally, for the block to run the thread must get the object lock first.

Typically, [Implementing the Runnable Interface|Runnable](implementing_the_runnable_interface|runnable) objects are used as their own locks, so that if you want to modify/access that object, you need its lock.



































#### Why:
#### How:









