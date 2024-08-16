# Proof That the Vertex Cover Approximation is a 2-approximation

Up: [Approximation Algorithms and Vertex Covers](approximation_algorithms_and_vertex_covers)
Brother(s):
TARGET DECK

Given our [Finding a Lower Bound for the Vertex Cover's Optimal Solution|lower bound](finding_a_lower_bound_for_the_vertex_cover's_optimal_solution|lower_bound), we know that for any set of disjoint edges the [Vertex Cover Algorithm](vertex_cover_algorithm) picks, the following is true:
$$ OPT(G) \ge |E^*| $$
Furthermore, the algorithm picks edges whose endpoints are not already in $C$, so we know that for every edge picked, 2 nodes are added to $C$:

$$|C| = 2|E^*|$$

Combining the equation above and the lower bound we get:
$$ \begin{align} |C| = 2|E^*|\\ \therefore |C| \le 2 \cdot OPT(G) \end{align} $$
And we are done.
































#### Why:
#### How:









