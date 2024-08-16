# Euclid's GCD with Recursion

Up: [Recursion](recursion)
Brother(s):
TARGET DECK

AKA Highest Common Factor, this is the highest integer $C$ for $A$ and $B$ that divides into both $A$ and $B$.

This is actually rather recursive, with its base case being $gcd(A, 0) = A$ and its recursive step being $gcd(A, B) = gcd(B, A\ mod\ B)$.

$gcd(80, 25) = gcd(25, 5) = gcd(5, 0) = 5$

[Implementation of gcd()](implementation_of_gcd())






































#### Why:
#### How:









