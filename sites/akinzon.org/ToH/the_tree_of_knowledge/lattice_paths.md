# Lattice Paths

Up: [Combinatorics](combinatorics)
Brother(s):
TARGET DECK

Lattice Paths in 2D is like moving on a grid of squares of even size.

In a 2D plane, the number of shortest lattice paths from $(0, 0)$ to $(m, n)$ is:
$$ \binom{m+n}{m} $$

$n$ can be used instead of m due to the symmetry of Pascal's triangle.

If going from $(a, b)$ to $(m, n)$, subtract $a$ and $b$ from both coordinates to give $(0, 0)$ to $(m-a, n-b)$.

![Pasted image 20240103151509.png](pasted_image_20240103151509.png)

****

### Lattice Paths and Trinomial Coefficient

In a 3D cube, the number of lattice paths from $(0, 0, 0)$ to $(m, n, p)$ is:
$$ \binom{m + n + p}{m,n,p} $$



> [!TIP]- How many lattice paths from $(0, 1, 2)$ to $(5, 5, 5)$ that pass through $(3, 3, 3)$ but not $(1, 2, 3)$?
> Do the question you nincompoop!


****

### Lattice Paths and Multinomial Coefficient

In a $n$-dimensional hypercube, the number of lattice paths from $(0, 0, ..., 0)$ to $(k1, k2, ..., kn)$ is
$$ \binom{k_1 + ... + k_n}{k_1,...,k_n} $$





























#### Why:
#### How:









