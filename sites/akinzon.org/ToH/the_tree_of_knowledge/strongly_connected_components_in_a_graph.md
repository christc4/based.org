# Strongly Connected Components in a Graph

Up: [Connected Components in a Graph](connected_components_in_a_graph)
Brother(s):
TARGET DECK

Strongly connected components occur only in directed [graphs](graphs), and are defined as graphs/subgraphs where there is a way to go from every vertex to another (but in directed graphs, since $(a, b)$ doesn't imply that there is an edge $(b, a)$).

The **Kosaraju** algorithm is used to find all strongly connected components:
	 - Call DFS on graph $G$ to compute finishing time for all nodes
	 - Create $G^T$ (Graph $G$ with all edges reversed)
	 - Call DFS in order of decreasing finishing time on $G^T$.































#### Why:
#### How:









