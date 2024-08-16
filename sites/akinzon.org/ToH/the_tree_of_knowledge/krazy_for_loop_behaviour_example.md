# Krazy For Loop behaviour example

Up: [For Loop](for_loop)
Brother(s): [Krazy For-each Loop behaviour example](krazy_for-each_loop_behaviour_example)
TARGET DECK

Consider the example:

```java
public class Main {  
  public static void main(String[] args) {  
    for (int i = 0; i < 9; i++); {  
      System.out.println("hello");  
    }  
  }  
}
```

> [!tip]- What will it do?
> This code will print 'hello' once, the code for printing hello is actually not inside the for loop due to the semicolon terminating the for loop before it could even begin

































#### Why:
#### How:









