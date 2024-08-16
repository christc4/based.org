# The Transient Keyword

Up: [Serialisation](serialisation)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::File IO

Yes, you're right. [The Externalizable Interface](the_externalizable_interface) is rather complicated to use.

using the keyword *transient* does the same thing, and is much simpler!

Let's say we wanted the name to be transient:

```java
private transient String name;
```

That does the trick! the `name` [Variables|variable](variables|variable) won't be serialised.

Why is the *transient* keyword better than implementing [the externalizable interface](the_externalizable_interface)? #flashcard 
They both allow you to choose what is serialised, but the transient keyword is simpler to use (no need to implement anything!)
<!--ID: 1707421397046-->


































#### Why:
#### How:









