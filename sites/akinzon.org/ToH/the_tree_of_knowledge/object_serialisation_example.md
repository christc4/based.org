# Object Serialisation Example

Up: [Object Serialisation](object_serialisation)
Brother(s): [Deserialisation Example](deserialisation_example)
TARGET DECK

`Person` class:

```java
import java.io.Serializable;

public class Person implements Serializable {
	private String name;
	private int age;

	public Person(String name, int age) {
		this.name = name;
		this.age = age;
	}

	public String toString() {
		return "Name: " + this.name + "; Age: " + this.age;
	}
}
```

`SerDemo1` class (for serialisation:)

```java
import java.io.File;
import java.io.FileOutputStream;
import java.io.ObjectOutputStream;
import java.io.OutputStream;

public class SerDemo1 {
	public static void main(String[] args) throws Exception{
		File f = new File("SerDemo1.txt");
		OutputStream out = new FileOutputStream(f);
		ObjectOutputStream oos = new ObjectOutputStream(out);
		oos.writeObject(new Person("Harry Potter", 18));
		oos.close();
	}
}
```

This above example will write all attributes of the `Person` class with name "Harry Potter" to a file.



































#### Why:
#### How:









