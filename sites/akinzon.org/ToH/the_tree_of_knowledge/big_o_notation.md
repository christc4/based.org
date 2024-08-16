# Big O Notation

Up: [Asymptotic Notation](asymptotic_notation)
Brother(s): [Big Theta Notation](big_theta_notation), [Big Omega Notation](big_omega_notation)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1201 Algorithmics::Asymptotic Notation

Big O notation gives an upper bound for the [running time of an algorithm](running_time_of_an_algorithm), if an algorithm had a big O of $O(n^2)$, it would execute **no more than** $n^2$ times.

a function $f(n)$ is a member of $O(g(n))$ if there exist constants $d \gt 0$ and $N \in \mathbb{N}$, such that:
$$ f(n) \le dg(n)\ for\ n \ge N $$
where $dg(n)$ acts as the upper bound.

Algorithms that are bounded by $O(n^2)$ are also bounded by $O(n^3)$ (it's just a higher ceiling, but not asymptotically tight?) .

An $O(n^2 )$ algorithm may not be faster than an $O(n^3 )$ algorithm, even for large $n$.

$c$ and $n_0$ must be independent of $n$.

**Why?** because that $O(n^3)$ algorithm could be an $O(n)$ algorithm in disguise.

[Big O Questions](big_o_questions)






























#### Why:
#### How:









