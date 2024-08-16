# Finding a Lower Bound for the Vertex Cover's Optimal Solution

Up: [Approximation Algorithms and Vertex Covers](approximation_algorithms_and_vertex_covers)
Brother(s):
TARGET DECK

The idea for the lower bound starts with the idea of disjoint edges in [Graphs](graphs), which are any two edges $(u,v)$ and $(x,y)$ where none of $u,v,x,y$ are the same.

For the lower bound, let $OPT(G)$ be the size of the minimal vertex cover for graph $G$.

For any set $E^* \subseteq E$ of pairwise disjoint edges we have:
$$ OPT(G) \ge |E^*| $$
Because the optimal solution will have vertices in it that will be in every edge in $E^*$, but may also have more vertices than can be found in $E^*$. In other words, if $E^* = E$, then $OPT(G) = |E^*|$. and you can replace the equal signs with $\subset$ and $\lt$ respectively.




































#### Why:
#### How:









