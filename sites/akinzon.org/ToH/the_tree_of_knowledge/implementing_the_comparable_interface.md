# Implementing the Comparable [Interfaces|Interface](interfaces|interface)

Up: [Implementation of Sets using Binary Search Trees](implementation_of_sets_using_binary_search_trees)
Brother(s):
TARGET DECK

It makes sense that to sort any node in our tree they must be comparable to each other, which simply requires us to implement the `Comparable` interface.

This is done by providing implementation for `x.compareTo(y)`, where `compareTo()` returns:

 - 0 if `x == y`
 - A negative number if `x < y`
 - A positive number if `x > y`

Being able to compare nodes comes in handy for [Finding Elements in a Set|finding elements](finding_elements_in_a_set|finding_elements)




































#### Why:
#### How:









