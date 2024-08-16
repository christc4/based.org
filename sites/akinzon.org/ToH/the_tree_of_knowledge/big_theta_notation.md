# Big $\Theta$ notation

Up: [Asymptotic Notation](asymptotic_notation)
Brother(s): [Big O Notation](big_o_notation), [Big Omega Notation](big_omega_notation)
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1201 Algorithmics::Asymptotic Notation


For a given function $g(n)$, $\Theta (g(n))$ is a [Sets|set](sets|set) of [functions](functions) where 
$$ \{f(n)\ |\ \exists\ c_1,c_2,n_0 \ge0\ \ s.t.\  0 \leq c_1g(n) \le f(n) \le c_2g(n)\  \forall n \ge n_0 \} $$
Where $n_0$ is the minimum value for which it holds, and $c_1$ and $c_2$ are any constant for which the inequality holds. There are infinitely many solutions, but you only need to find 1 set of em.

That's quite a lot to look at, but what it means is that $\Theta (g(n))$ is a set of functions that lie between the worst and best [Running Time of an Algorithm|running times](running_time_of_an_algorithm|running_times).

All members of $\Theta (g(n))$ grow at the same rate.

[Advantages of big Theta notation](advantages_of_big_theta_notation)

[Disadvantages of big Theta notation](disadvantages_of_big_theta_notation)

If $f(n) = O(g(n))$ (equals really means set membership) and $f(n) = Ω(g(n))$, then $f(n)$ is constrained both by the algorithm's upper and lower bound, and is candid to be a member of $\Theta (g(n))$.

Saying that the running time of an algorithm is $\Theta(n)$ is a strong statement, because it means that it is both $O(n)$ in the worst case and $\Omega (n)$ in the best case.

https://www.khanacademy.org/computing/computer-science/algorithms/asymptotic-notation/a/big-big-theta-notation

[Big Theta Questions](big_theta_questions)




























#### Why:
#### How:









