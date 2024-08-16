# Breadth-First Time Complexity

Up: [Breadth-First Search](breadth-first_search)
Brother(s): [Time Complexity of Depth-First Search](time_complexity_of_depth-first_search)
TARGET DECK

Since no node is visited more than once, breadth-first search is equivalent to enqueueing and dequeueing each node exactly once, and dequeueing is $O(1)$. 

We do this $V$ times, so it is $O(|V|)$.

We also perform a bunch of operations for each neighbour of a node, this value can be called $x$. doing $x$ operations $V$ times is $O(|E|)$.

Overall, the time complexity is $O(|V| + |E|)$.


































#### Why:
#### How:









