# Character-Based Writer Stream Example

Up: [Character-Based Streams](character-based_streams)
Brother(s):
TARGET DECK

```java
import java.io.File;
import java.io.FileWriter;
import java.io.Writer;

public class FileDemo{
	public static void main(String[] args) throws Exception{
		//Step 1: ready the file
		File f = new File("d:" + File.separator + "test.txt");
		//Step 2: set up the stream
		Writer out = new FileWriter(f);
		//Step 3: start writing
		String str = "Hello World!";
		out.write(b);
		//Step 4: close the stream
		out.close();
	}
}
```

Adding `true` to the `FileWriter`'s constructor has the same effect as it does on `FileOutputStream`.



































#### Why:
#### How:









