# Detecting Bipartite Graphs

Up: [Bipartite Graphs](bipartite_graphs)
Brother(s):
TARGET DECK

Bipartite Graphs can be detected using a modified [Breadth-First Search](breadth-first_search).

Since the breadth-first search "colours" in nodes once discovered, we can use this to our advantage!

Colour in the source node your primary colour, lets say green.

Then colour in all its neighbours the secondary colour, lets say red.

If as you search the tree you find that the neighbour of a node has the same colour as it, the tree is not bipartite.



































#### Why:
#### How:









