# Replacing [Basic Variables](basic_variables) with [Non-Basic Variables](non-basic_variables)

Up: [The Simplex Method](the_simplex_method)
Brother(s):
TARGET DECK

To replace a basic variable with a non-basic variable, this is what happens:
	- Choose a non-basic variable in the [the objective function|objective function](the_objective_function|objective_function) whose coefficient is positive (meaning they can still contribute to maximising $z$)
	- Calculate the maximum value it can take in all of the basic variable equations, such that $b_i - a_mx_j$ where $a_m$ is the coefficient of the non-basic variable $x_j$ in the basic variable equation is positive
	- The basic variable equation in which $x_j$ has a minimal value is then written in terms of $x_j$ 
	- All occurrences of $x_j$ are replaced with this expression.



































#### Why:
#### How:









