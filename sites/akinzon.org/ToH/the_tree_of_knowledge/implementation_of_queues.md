# Implementation of Queues

Up: [Queues](queues)
Brother(s):
TARGET DECK

When implemented with [linked lists](linked_lists), queue operations take $\Theta (1)$ time, since:
	- `enqueue()`ing at the `tail` doesn't rely on the queue
	- `dequeue()`ing at the `head` doesn't rely on the queue

For enqueueing at the tail, it is only $\Theta (1)$ if we use a [Doubly Linked Lists|Doubly Linked List](doubly_linked_lists|doubly_linked_list), because otherwise we would need to look for the new tail, because in [Singly Linked Lists](singly_linked_lists), nodes do not know which node comes before them, only the one that comes after.




































#### Why:
#### How:









