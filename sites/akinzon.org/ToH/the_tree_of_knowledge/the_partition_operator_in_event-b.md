# The Partition Operator in Event-B

Up: [Partitions in Event-B](partitions_in_event-b)
Brother(s):
TARGET DECK

The `partition()` operator is a shorthand way of declaring that a [Sets|set](sets|set) is to be *partitioned* into disjoint subsets.

**Usage: `partition(light, on, off)`**

Lights are either on or off.

> [!tip] Using subsets of a [Carrier Sets|carrier set](carrier_sets|carrier_set) to represent states allows us to easily define what can be done if the item is in one of a number of specific states! 
> For example, if a printer had the states `on`, `off`, `busy`, if it was `off` or `busy`, it would tell the user that the printer cannot be used at the moment, which can be represented by a simple set union!
> You couldn't do this if the sets weren't partitions..































#### Why:
#### How:









