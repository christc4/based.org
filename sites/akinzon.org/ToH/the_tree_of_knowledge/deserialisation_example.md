# Deserialisation Example

Up: [Deserialisation](deserialisation)
Brother(s): [Object Serialisation Example](object_serialisation_example)
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

`SerDemo2` class (for deserialisation:)

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.ObjectInputStream;
import java.io.InputStream;

public class SerDemo2 {
	public static void main(String[] args) throws Exception{
		File f = new File("SerDemo1.txt");
		InputStream input = new FileInputStream(f);
		ObjectInputStream ois = new ObjectInputStream(input);
		Object obj = ois.readObject();
		ois.close();
		System.out.println(obj);
	}
}
```

With the output being the object's attributes in the programmer-defined `toString()` format.



































#### Why:
#### How:









