# Vertex Cover Algorithm

Up: [Approximation Algorithms and Vertex Covers](approximation_algorithms_and_vertex_covers)
Brother(s):
TARGET DECK

The algorithm works by starting at an arbitrary edge $(u,v)$ in $E$, and adding both $u$ and $v$ to $C$, then discarding all other edges containing either $u$ or $v$ from $E$, and continuing until there are no more edges left.

```
C = Set.emptySet
E_c = E
while E_c != Set.emptySet {
	select edge (u,v) in E_c
	C = C U {u,v}
	remove all edges incident to u and v from E_c
}
return C
```




































#### Why:
#### How:









