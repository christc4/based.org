# Krazy For-each Loop behaviour example

Up: [For each (enhanced for) Loop](for_each_(enhanced_for)_loop)
Brother(s): [Krazy For Loop behaviour example](krazy_for_loop_behaviour_example)
TARGET DECK

Consider this code:

```java
public class Main {  
  public static void main(String[] args) {  
    int[] numbers = new int[2];  
    numbers[0] = 1;  
    numbers[1] = 1;  
    for (int num: numbers){}; {  
      System.out.println(num);  
    }  
  }  
}
```

> [!tip]- What will it do?
> This code will not compile, printing `num` will not work since it not within the scope of the for loop where `num` is defined.



































#### Why:
#### How:









