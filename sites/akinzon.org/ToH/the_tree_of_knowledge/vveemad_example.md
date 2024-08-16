# VVEEMAD example

Up: [VVEEMAD](vveemad)
Brother(s):
TARGET DECK

Consider this graph:
	![Pasted image 20240103210636.png](pasted_image_20240103210636.png)

Its $V$ would be $\{a, b, c, d\}$.

Its $|V|$ would be $|\{a, b, c, d\}|$ which is 4.

Its $E$ would be $\{(a, b), (a, c), (a, d), (b, c), (b, d), (c, d)\}$.

Its $|E|$ would be 6.

Its $M$ would be:
$$ \begin{bmatrix}  
0 & 1 & 1 & 1\\  
1 & 0 & 1 & 1\\
1 & 1 & 0 & 1\\
1 & 1 & 1 & 0
\end{bmatrix}$$
And due to the nature of this graph (undirected), its $A$ would be the same.

An [Adjacency Lists|adjacency list](adjacency_lists|adjacency_list) for this graph would be:

$a \rightarrow b, c, d$
$b \rightarrow a,c,d$
$c \rightarrow a,b,d$
$d \rightarrow, a,b,c$

It's $D$ would be:
$$ \begin{bmatrix}  
3 &  &  & \\  
 & 3 &  & \\
 &  & 3 & \\
 &  &  & 3
\end{bmatrix}$$






























#### Why:
#### How:









