# Krazy Iterator casting example

Up: [Iterators](iterators)
Brother(s):
TARGET DECK

Consider this code:

```java
public class Main {  
  public static void main(String[] args) {  
    ArrayList<Dog> dogArrayList = new ArrayList<Dog>();  
    for (int i = 0; i < 10; i++){  
      dogArrayList.add(new Dog("Peanut" + i));  
    }  
    Iterator it = new dogArrayList.iterator();  
  
    while (it.hasNext()) {  
      Dog dog = it.next();  
      System.out.println(dog.getName());  
    }  
  }  
}
```

> [!tip]- Would it run?
> This code would not compile because the object returned by `it.next()` would be of type `java.lang.Object` and not the necessary `Dog` type. To fix this, the value would have to be casted to `Dog` by prepending `(Dog)` before the call to `it.next()`.



































#### Why:
#### How:









