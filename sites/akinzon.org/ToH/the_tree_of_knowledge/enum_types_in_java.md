# Enum Types in Java

Up: [Enum Types](enum_types)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::Enum Types

In Java, Enum Types are actually an example of anonymous inner classes.

This:

```Java
public enum Country{ENGLAND, NIGERIA, BRAZIL}
```

Becomes this:

```Java
static public abstract class Country {
	private Country() {}; // actual declaration of Country class
	public static final Country ENGLAND = new Country();
	public static final Country NIGERIA = new Country();
	public static final Country BRAZIL = new Country();
}
```
































#### Why:
#### How:









