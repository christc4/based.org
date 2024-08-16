# Nullspace of a Matrix

Up: [Matrices](matrices)
Brother(s):
TARGET DECK

The Kernel (or **Nullspace**) of a matrix is the [Sets|set](sets|set) of all non-trivial solutions to $$ Ax = 0 $$

p n-dimensional vectors are linearly dependent if n < p.
This is to say that you cannot have more than n independent vectors in an n-dimensional space.

### Calculating the nullspace.

Classic trick to finding the Nullspace of a matrix:
https://www.khanacademy.org/math/linear-algebra/vectors-and-spaces/null-column-space/v/dimension-of-the-null-space-or-nullity


1) Find reduced row echelon form for matrix since its easier to calculate nullspace that way
2) Determine the pivots (leading 1's in a row, only 1 in column) and free variables (non-pivots)
3) Determine the value of the pivots by viewing the matrix as a linear system of equations
4) The nullspace will be a linear combination of each free variable $\mathcal{x}_i$ multiplied by some vector with $n$ entries. Where $i = n$, the entry is 1. In the same $n$ for a different $x_i$, the entry is 0.
5) The pivots ($\bar x_ \bar i$) will give you coefficients on the free variables when you solve for the pivots, these are entries on the $\bar i^{th}$ row. 
6) If linearly independent, they can be used as a basis for the nullspace.
7) The dimension of the nullspace is the number of free variables.
































#### Why:
#### How:









































#### Why:
#### How:









