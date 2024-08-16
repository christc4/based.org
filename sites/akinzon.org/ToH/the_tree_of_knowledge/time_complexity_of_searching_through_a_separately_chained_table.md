# Time Complexity of Searching Through a Separately Chained Table

Up: [Separate Chaining](separate_chaining)
Brother(s):
TARGET DECK

For a separately chained table, if the values are spread out, search and insertion both take $O(1)$ time.

If there are many values that hash to the same address, the table becomes a glorified [Linked Lists|linked list](linked_lists|linked_list), giving a search time of $O(n)$.

But generally, if your [Hashing Algorithms|hashing algorithm](hashing_algorithms|hashing_algorithm) is good and the hash table is of a good size, you can expect $\Theta (1)$ performance.



































#### Why:
#### How:









