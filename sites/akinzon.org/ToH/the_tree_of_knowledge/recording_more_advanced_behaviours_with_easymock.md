# Recording More Advanced Behaviours with EasyMock

Up: [EasyMock](easymock)
Brother(s):
TARGET DECK

The `expect()` [Methods|method](methods|method) can be used to specify other things for the mock class to do

Call a method $n$ times:
```java
expect(mock.someMethod()).times(n);
```


Specify return values:
```java
expect(mock.someMethod()).andReturn("shuush");
```

Respond with an [Exceptions|exception](exceptions|exception):
```java
expect(mock.someMethod()).andThrow(new IndexOutOfBoundsException());
```

Chaining them together:
```java
expect( mock.someMethod() ) 
	.andReturn("Banana") 
	.times(20) 
	.andThrow( new ArithmeticException() ) 
	.andReturn("No more bananas left!");
```





























#### Why:
#### How:









