# Modelling Stacks in Event-B

Up: [Representing Collections in Event-B](representing_collections_in_event-b)
Related Note(s): [Stacks](stacks)
TARGET DECK

As you know, stacks are a LIFO abstract data type where the order of insertion is the reverse of the order of removal, the last into the stack is the first one out.

In Event-B, we can view a stack as a LIFO [Modelling Queues in Event-B|queue](modelling_queues_in_event-b|queue), which in turn can be viewed as a [Modelling Lists in Event-B|list](modelling_lists_in_event-b|list) with some extra rules.

Well for a stack, we slightly reverse those rules!

There is still a [Total Functions in Event-B|total](total_functions_in_event-b|total) [Functions (Advanced)|injective](functions_(advanced)|injective) function from each item to its position, but here, we remove the item with the highest position, and the item we add also has the new highest position too.



































#### Why:
#### How:









