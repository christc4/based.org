# Application of Depth-First Search - Topological Sorting

Up: [Depth-First Search](depth-first_search)
Brother(s):
TARGET DECK

Topological Sorting is a way of ordering nodes in a [Graphs|graph](graphs|graph) such that all "dependencies" of a node are to its left in the order.

More formally, for every edge $(u, v)$ in a graph from node $u$ to node $v$, **$u$ always comes before $v$.**

Topological sorting can be done using DFS:
	- Call DFS on a graph
	- Every time a node is finished being processed (popped of the [Stacks|stack](stacks|stack)) it is added to the front of a [Linked Lists|linked list](linked_lists|linked_list)
	- The resulting list at the end of the program is the topological sorting.





































#### Why:
#### How:









