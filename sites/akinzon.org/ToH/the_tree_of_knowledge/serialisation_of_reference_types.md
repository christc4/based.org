# Serialisation of Reference Types

Up: [Serialisation](serialisation)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1206 Programming II::File IO

Say, for example, we [Object Serialisation|serialise an object](object_serialisation|serialise_an_object). This means that its member variables are stored in a file, along with the [Classes|class](classes|class) name. But what if those member variables are actually [objects](objects) themselves?

Well, [Recursion|recursion](recursion|recursion) is our best friend here, it can go ahead and serialise that object too! So long as it also implements the *Serializable* [Interfaces|interface](interfaces|interface) as well. (well if it didn't, we'd either have to declare it transient or the object that it belongs to is not serialisable.)


If an object has been serialised, does it get serialised again? #flashcard 
No, [You Only Get Serialised Once](you_only_get_serialised_once).
<!--ID: 1707421214307-->



[You Only Get Serialised Once](you_only_get_serialised_once)
































#### Why:
#### How:









