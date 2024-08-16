# The Importance of a Dummy Marker for Deletion

Up: [Hash Tables](hash_tables)
Brother(s):
TARGET DECK

To delete a value in a hash table, we must first search for it.

If we cannot find it in the hashed-to address due to a [Hash Collisions|collision](hash_collisions|collision), we must [Iterating Through a Hash Table|iterate through the hash table](iterating_through_a_hash_table|iterate_through_the_hash_table) in a [linear probing](linear_probing) fashion until we find the value or a `null` entry (assuming that the value was there and was deleted, since why shouldn't it have occurred any free space after the address it couldn't fit in?).

However, this could be a wrong assumption, perhaps it went further than where the `null` value is, but now we won't get there!

To distinguish between empty spaces and spaces that had something deleted from them, a dummy marker containing a value that **cannot normally be in the hash table** will help us in many ways:
	- The `find()` method won't see it as an empty space
	- Iterators can ignore dummy entries
	- the `insert()` method can insert in dummy locations




































#### Why:
#### How:









