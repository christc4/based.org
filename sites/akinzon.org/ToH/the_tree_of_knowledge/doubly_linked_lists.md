# Doubly Linked Lists

Up: [Linked Lists](linked_lists)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1201 Algorithmics::Data Structures

Unlike [Singly Linked Lists](singly_linked_lists), doubly linked lists also store a pointer to the previous node in each node.

They can also use dummy nodes that help with implementation (for circular lists, or checking empty lists.)

[Time Complexity of Doubly Linked Lists](time_complexity_of_doubly_linked_lists)

What is the benefit of a dummy (sentinel) node in [Linked Lists](linked_lists)? #flashcard 
They help with implementing [Doubly Linked Lists](doubly_linked_lists), we no longer need to store a separate `head` and `tail` reference. We can easily check if a list is empty, and adding to the head and tail is easier too.
Furthermore, it removes the need to check if the next node is `null` since the last node will just point to the sentinel node.
<!--ID: 1714327802115-->




































#### Why:
#### How:









