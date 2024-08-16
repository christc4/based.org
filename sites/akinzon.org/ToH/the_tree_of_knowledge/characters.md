# Characters

Up: [Primitives](primitives)
Brother(s):
TARGET DECK

Characters are a type of primitive that are used to store different characters according to some standard character set, namely, 16-bit Unicode.

In Java, characters are surrounded by single quotes.

For example: 

```java
char c = 'z';
```

The actual value stored in `c` is not the literal 'z', but the 16-bit [Unicode](unicode) value (which would be u007A in hex.)

Special characters must be drawn using `new ShowString(char)` which uses `Graphics.drawString()` since they cannot simply be keyed in from the keyboard.

[Characters and Character-based Streams](characters_and_character-based_streams)

[Strings](strings) are a sequence of characters.































#### Why:
#### How:









