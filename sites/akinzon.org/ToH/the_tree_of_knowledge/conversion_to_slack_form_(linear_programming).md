# Conversion to Slack Form (Linear Programming)

Up: [Slack Form for Linear Programming](slack_form_for_linear_programming)
Brother(s):
TARGET DECK

Conversions from [Standard Form for Linear Programming|standard form](standard_form_for_linear_programming|standard_form) to slack from is therefore done by turning those linear inequalities into linear equalities.

This is done by introducing **[slack variables](slack_variables).**

For each [Linear Programming Constraints|constraint](linear_programming_constraints|constraint) $\sum_{j=1}^{n} a_{ij}x_j \le b_i$, a slack variable $s_i$ is introduced where $\sum_{j=1}^{n} a_{ij}x_j + s_i = b_i$ and $s_i \ge 0$.

Slack variable get mixed with the original variables during the [The Simplex Method|simplex method](the_simplex_method|simplex_method) and are renamed to $x_{n+i}$.


































#### Why:
#### How:









