# Byte-Based Input Stream Example

Up: [Byte-Based Streams](byte-based_streams)
Brother(s):
TARGET DECK

```java
import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;

public class FileDemo{
	public static void main(String[] args) throws Exception{
		//Step 1: ready the file
		File f = new File("d:" + File.separator + "test.txt");
		//Step 2: set up the stream
		InputStream input = new FileInputStream(f);
		//Step 3: start reading
		byte b[] = new byte[(int)f.length()];
		input.read(b);
		//Step 4: close the stream
		input.close();
		System.out.println(new String(b));
	}
}
```

**Important to set the size of a bytearray as the file size**



































#### Why:
#### How:









