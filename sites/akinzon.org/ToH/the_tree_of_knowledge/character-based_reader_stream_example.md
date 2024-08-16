# Character-Based Reader Stream Example

Up: [Character-Based Streams](character-based_streams)
Brother(s):
TARGET DECK
```java
import java.io.File;
import java.io.FileReader;
import java.io.Reader;

public class FileDemo{
	public static void main(String[] args) throws Exception{
		//Step 1: ready the file
		File f = new File("d:" + File.separator + "test.txt");
		//Step 2: set up the stream
		Reader reader = new FileReader(f);
		//Step 3: start reading 
		int temp = 0;
		while ((temp = reader.read())!= -1) {
			System.out.println((char)temp);
		}
		//Step 4: close the stream
		reader.close();
	}
}
```

`reader.read()` - returns the read character, or `-1` to signify end-of-stream.







































#### Why:
#### How:









