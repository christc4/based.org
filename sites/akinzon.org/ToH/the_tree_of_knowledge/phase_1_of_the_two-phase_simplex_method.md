# Phase 1 of the Two-Phase Simplex Method

Up: [Two-Phase Simplex Method](two-phase_simplex_method)
Brother(s):
TARGET DECK

A new variable $x_0$ is introduced, along with its auxiliary problem:

$$  \begin{align} 

\text{maximise } - x_0\\
\text{subject to} \sum_{j=1}^{n} a_{ij}x_j - x_0 \le b_i, \text{ for } i = 1...m\\
\text{where } x_j \ge 0, \text{ for } j = 1...n\\
\end{align}  $$

The original problem is feasible **if and only if** the auxiliary problem has a solution where $x_0 = 0$.

When solving the auxiliary problem, the first [Pivot Operations (Linear Programming)|pivot operation](pivot_operations_(linear_programming)|pivot_operation) **always** chooses $x_0$ as the entering variable.

[Example of Creating an Auxiliary Problem from a Base Problem](example_of_creating_an_auxiliary_problem_from_a_base_problem)




































#### Why:
#### How:









