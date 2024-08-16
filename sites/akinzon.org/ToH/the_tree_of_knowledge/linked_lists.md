# Linked Lists

Up: [List of Some Data Structures](list_of_some_data_structures)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1201 Algorithmics::Data Structures

Linked lists are a type of data structure, comprised of nodes.

These nodes can be described by this definition:

```java
class Node<T>
{
	private T element;
	private Node<T> next;
}
```

where `element` is the actual value of the node,

and `next` is the pointer to the subsequent node.

Linked lists have a `head` which points to the first node,
and sometimes a `tail` that points to the last node.

Relevant linked list operations:
	 - Create a linked list
	 - [Adding to a Linked List|Add elements](adding_to_a_linked_list|add_elements)
	 - [Removing from a Linked List|Remove elements](removing_from_a_linked_list|remove_elements)
	 - [Checking if Element Exists in Linked List|Check if element exists](checking_if_element_exists_in_linked_list|check_if_element_exists)
	 - From the above note, we see that any operation that requires going through the list is $\Theta (n)$. Linked lists aren't that great when needing to go through the list.

Linked lists come in many forms:

[Singly Linked Lists](singly_linked_lists)

[Doubly Linked Lists](doubly_linked_lists)

[When To Use Linked Lists](when_to_use_linked_lists)

[Skip Lists](skip_lists)

What are [linked lists](linked_lists)? #flashcard 
Linked lists are a data structure that store data as nodes, which store a reference to the next node (and previous node for [doubly linked lists](doubly_linked_lists))














































