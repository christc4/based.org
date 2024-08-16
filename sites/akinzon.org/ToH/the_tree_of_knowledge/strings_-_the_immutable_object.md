# Strings - The Immutable Object

Up: [Strings](strings)
Brother(s):
TARGET DECK

String [Objects](objects) are **immutable** - this means they cannot be changed.

#### B-but muh concatenation!!!
When concatenating a string to another, the contents of the string aren't simply changed.

Instead, this process happens, for adding the contents of a string literal to `strA`:
1) The contents of the string literal are added to the [The String Pool|string pool](the_string_pool|string_pool).
2) This is then added with the contents of `strA` and stored in another location in the string pool. This is the final concatenation result
3) `strA`'s reference to its original value is changed to point to the final concatenation result.

#### The immutable object meets the unstoppable modifier

using `java.lang.StringBuffer` and append() can also get around the immutability of strings.


































#### Why:
#### How:









