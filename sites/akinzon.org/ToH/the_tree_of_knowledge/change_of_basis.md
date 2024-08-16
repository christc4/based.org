# Change of Basis

Up: [Linear Algebra](linear_algebra)
Brother(s):
TARGET DECK

The thing about [matrices|matrix transformations](matrices|matrix_transformations) in general is that they are all in the common tongue, the standard language, the basis $I$  that is $$\begin{bmatrix} 1\  0 \\ 0\ 1\end{bmatrix}$$ But what if we wanted to represent these transformations in coordinate system with a different basis $A$ ?, let's say: $$\begin{bmatrix} 1\ 2 \\ 1\ 3\end{bmatrix}$$
??


### Eu Falo Portugues
For a vector $$ \begin{bmatrix} -1 \\ 2 \end{bmatrix} $$ This may be how it is represented in our language, but a 
$\begin{bmatrix} -1 \\ 2 \end{bmatrix}$ in the language of basis $A$ would look entirely different. It would actually be $\begin{bmatrix} 3 \\ 5 \end{bmatrix}$ in basis $A$! 

To understand how to see our vectors in their basis, recall that vectors are merely linear combinations of the bases of coordinate systems, and each component of a vector scales the corresponding column of the basis. Instead of automatically scaling these by the basis $I$, we would scale them by basis $A$, giving us the result.

To see their vectors in our basis, we would need the inverse of their basis, and apply that to their vector.
### Improvise, adapt, overcome
We would need to first figure out where our basis [vectors](vectors) land, which is kindly given to us by the matrix above. Now for a transformation $$\begin{bmatrix} 3\ 4 \\ 2\ 3\end{bmatrix}$$
We apply it to the matrix that converts basis $A$ into our language of basis $I$ , then transform back using $A^{-1}$. 
































#### Why:
#### How:









