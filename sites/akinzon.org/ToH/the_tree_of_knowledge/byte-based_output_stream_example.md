# Byte-Based Output Stream Example

Up: [Byte-Based Streams](byte-based_streams)
Brother(s):
TARGET DECK

```java
import java.io.File;
import java.io.FileOutputStream;
import java.io.OutputStream;
import java.io.DataOutputStream;

public class FileDemo{
	public static void main(String[] args) throws Exception{
		//Step 1: ready the file
		File f = new File("d:" + File.separator + "test.txt");
		//Step 2: set up the stream
		OutputStream out = new FileOutputStream(f);
		//Step 3: start writing
		String str = "Hello World!";
		OutputStream dos = new DataOutputStream(out);
		dos.writeChars(str);
		//Step 4: close the stream
		out.close();
	}
}
```

**NOTE:** Files are typically overwritten, but adding `true` to the constructor on the `FileOutputStream` causes new content to be appended to the end of the file.

Newline character - `\r\n`



































#### Why:
#### How:









