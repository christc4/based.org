# The Motivation for Hash Tables

Up: [Hash Tables](hash_tables)
Brother(s):
TARGET DECK

Imagine we had a list of object, and we wanted to have look-ups that allow us to find the corresponding object.

This is known as **content-addressable memory.**

We could use a [lists|list](lists|list), but finding something in a list is $\Theta (n)$. 

We could use [trees](trees) to give us $\Theta(log(n))$ but that only works if the data is sorted!

We could just use a balanced [Binary Search Trees|binary search tree](binary_search_trees|binary_search_tree) to keep it ordered though.

Or we could use something that stores a 'key' for each 'value', in pairs. Given the key, wouldn't this even be $O(1)?$

**This is the power of the Hash Table.**

































#### Why:
#### How:









