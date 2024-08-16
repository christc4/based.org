# Branch-And-Bound Integer Linear Programming

Up: [Branch-And-Bound](branch-and-bound)
Related Note(s): [Integer Linear Programming](integer_linear_programming)
TARGET DECK

Formulating integer linear programming problems requires **relaxation** on the integer constraints so that we can use regular linear programming.

Getting the best solution from there requires branch-and-bound to get the best possible set of values.

By relaxing the constraints we obtain a larger feasible area, and a more relaxed solution $\bar z^*$.

Compared to the actual integer linear programming solution $z^*$, $\bar z^* \le z^*$.

We solve a new LP problem at each branch, and act according to these cases:
	- If the solution of the current LP problem ($z$) $\ge z^*$ (best found value) or if the problem is infeasible we can't get any better, so we backtrack.
	- If $z \lt z^*$:
		- if all the values ([Basic Variables](basic_variables) and [Non-Basic Variables](non-basic_variables)) are integers then we set $z^* = z$ and backtrack.
		- if not, we branch by setting each non-integer $x$ to either $\lfloor x \rfloor$ or $\lceil x \rceil$. 

































#### Why:
#### How:









