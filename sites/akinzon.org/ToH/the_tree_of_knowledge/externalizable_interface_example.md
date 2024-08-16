# Externalizable Interface Example

Up: [The Externalizable Interface](the_externalizable_interface)
Brother(s):
TARGET DECK

`Person` class:

```java
import java.io.Externalizable;  
import java.io.IOException;  
import java.io.ObjectInput;  
import java.io.ObjectOutput;  
  
public class Person implements Externalizable {  
  private String name;  
  private int age;  
    
  public Person() {}  
    
  public Person(String name, int age) {  
    this.name = name;  
    this.age = age;  
  }  
    
  public String toString(){  
    return "Name: " + this.name + "; Age: " + this.age;  
  }  
  
  @Override  
  public void writeExternal(ObjectOutput out) throws IOException {  
    out.writeObject(this.name);  
    out.writeInt(this.age);  
  }  
  
  @Override  
  public void readExternal(ObjectInput in) throws IOException, ClassNotFoundException {  
    this.name = (String)in.readObject();  
    this.age = in.readInt();  
  
  }  
}
```

`SerDemo3` class 

```java
import java.io.*;  
  
public class SerDemo3 {  
  public static void main(String[] args) throws Exception{  
    ser();  
    dser();  
  }  
  
  public static void ser() throws Exception{  
    File f = new File("SerDemo3.txt");  
    OutputStream out = new FileOutputStream(f);  
    ObjectOutputStream oos = new ObjectOutputStream(out);  
    oos.writeObject(new Person("Ron", 17));  
    oos.close();  
  }  
  
  public static void dser() throws Exception{  
    File f = new File("SerDemo3.txt");  
    InputStream input = new FileInputStream(f);  
    ObjectInputStream ois = new ObjectInputStream(input);  
    Object obj = ois.readObject();  
    ois.close();  
    System.out.println(obj);  
  }  
}
```

`readExternal` and `writeExternal` are used to filter between which attributes are serialised.
































#### Why:
#### How:









