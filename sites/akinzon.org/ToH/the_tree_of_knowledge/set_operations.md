# Set Operations

Up: [Sets](sets)
Brother(s):
TARGET DECK

#### Union
A ∪ B: Union - Combining two sets and eliminating any repetitions
$|A ∪ B| = |A| + |B| - |A ∩ B|$
(Associative)
(Commutative)

#### Sum, or Disjoint Union
A + B: Disjoint Union - A special case of union where repetitions are NOT eliminated because each member of each set is marked out as belonging to that set, for example:
$X+Y = \{ (x,0) | x∈X\} ∪ \{(y,1) | y∈Y\}$
^ Even if X and Y had an intersection, lets say "z", set X's version of z becomes (z, 0) and set Y's is $(z, 1)$, and $(z, 0) ≠ (z, 1)$
$|A + B| = |A| + |B|$
$|A + B| = |A ∪ B|$ **if and only if** sets A and B are **disjoint** (never had any intersection to begin with)

#### Intersection
A ∩ B: Intersection - Only counting elements belonging to both A AND B.
$|A ∩ B| = |A| + |B| - |A ∪ B|$
(Associative)
(Commutative)

#### Set Difference
A - B: Difference - Removing the intersection of a set A with set B from set A 
$|A - B| = |A| - |A ∩ B|$
(Not commutative)
(Not associative)

#### The Cartesian Product
A x B: Cartesian Product - The set of all ordered pairs that can be created from two sets A and B
$X×Y = \{ (x,y) | x∈X\ \&\ y∈Y\}$
$|A \times B| = |A| \times |B|$
(Not commutative, see [Ordered Pairs](ordered_pairs))

#### THE POWERSET.
$\mathcal{P}(A)$: The powerset of set A: This is the set of all subsets of A, including A itself (by X⊆X) and The Empty Set.
$|\mathcal{P}(A)| = 2^{|A|}$


#### [Functions](functions) from one set to another
$Y^X$: This is the set of all functions from set $X$ to set $Y$. This also encodes the **function space** between them.
$|Y^X| = |Y|^{|X|}$ 
























#### Why:
#### How:









