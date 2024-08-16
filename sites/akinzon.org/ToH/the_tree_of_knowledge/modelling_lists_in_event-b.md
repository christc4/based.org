# Modelling Lists in Event-B

Up: [Representing Collections in Event-B](representing_collections_in_event-b)
Related Note(s): [Lists](lists)
TARGET DECK

As you should know, lists are an abstract data type in which data is held in a particular order, and allows for repeated values.

At its most abstract, a list is a bunch of **items** that each have a **distinct** **position** in the list.

So what would be the best way to relate an item to its position?

(hehe) **Relations!!!** 

Yes, if we create a mapping $position:\ item \rightarrow POSITION$, we can call $position(item)$ to return that item's position.

But which type of mapping will we need? There are so many, and we need to make sure that each item's position is distinct!

**Injective [Functions (Advanced)|Functions](functions_(advanced)|functions)** are the key, they will only let each item have a 1:1 mapping with any position, so no two items have the same position.

Furthermore, we would like to make this mapping a [Total Functions in Event-B|total function](total_functions_in_event-b|total_function), since this will assure that **every item has a position.**




































#### Why:
#### How:









