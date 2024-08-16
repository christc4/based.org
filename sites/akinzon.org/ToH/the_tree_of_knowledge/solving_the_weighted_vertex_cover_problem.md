# Solving The Weighted Vertex Cover Problem

Up: [Weighted Vertex Covers](weighted_vertex_covers)
Brother(s):
TARGET DECK

This problem is best solved using [Linear Programming](linear_programming) rather than with any [Vertex Cover Algorithm](vertex_cover_algorithm).

Each vertex $v \in V$ is represented by **decision variables** $x_v$.

These variables equal 1 if $v \in C$ and 0 otherwise.

As per the [Weighted Vertex Cover Problem Background|problem background](weighted_vertex_cover_problem_background|problem_background), we want to minimise the total weight of the vertices in the vertex cover.

The [The Objective Function|objective function](the_objective_function|objective_function) to minimise is:

$$ z = \sum_{v \in V} w(v) \cdot x_v $$

Of course, we'll need some [Linear Programming Constraints|constraints](linear_programming_constraints|constraints):

Intuitively, the constraint is that for any edge $(u,v) \in E$:

$$ (x_u = 1) \lor (x_v = 1) $$
However, these aren't in a linear form. So we put them in a linear form like so:

$$  x_u + x_v \ge 1, \forall(u,v) \in E $$

Which is the same because no matter what, so long as either $x_u$ or $x_v$ is 1, the condition is satisfied.

The other constraint is that 

$$x_v \in \{0,1\}, \forall v \in V$$
To represent this in a linear programming way, we **relax** the constraints to $0 \le x_v \le 1$, which is less strict but is linear.

Furthermore, during the process of calculating the linear programming solution, we may get values for $x_v$ that are not 0 or 1. We use a rounding rule in that situation:

$$ x_v = \begin{cases} 1, \text{if } x_v \ge 0.5 
\\ 0, \text{otherwise}
\end{cases} $$

This is because the constraint $x_u + x_v \ge 1$ implies that either $x_u \ge 0.5$ or $x_v \ge 0.5$ or both, and so at least one of $x_u$ and $x_v$ is in $C$ for every edge $(u, v) \in E$.

The final [Linear Programming Problems|problem](linear_programming_problems|problem) is as follows:

$$ \begin{align} 

\text{minimise } z = \sum_{v \in V} w(v) \cdot x_v\\

\text{subject to  } x_u + x_v \ge 1\ \ \ \ \ \ \ \ \forall(u,v) \in E\\
x_v \le 1\ \ \ \ \ \ \ \ \ \ \forall v\in V\\
x_v \ge 0 \ \ \ \ \ \ \ \ \ \ \forall v\in V
\end{align}$$
























#### Why:
#### How:









