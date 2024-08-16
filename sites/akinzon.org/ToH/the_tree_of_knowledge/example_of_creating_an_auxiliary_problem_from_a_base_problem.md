# Example of Creating an Auxiliary Problem from a Base Problem

Up: [Phase 1 of the Two-Phase Simplex Method](phase_1_of_the_two-phase_simplex_method)
Brother(s):
TARGET DECK

**This:**

$$
 \begin{align} 

\text{maximise } z = 2x_{1} - x_2\\
\text{subject to } 2x_1 - x_2 \le 2\\
x_1 - 5x_2 \le -4\\
\text{where } x_1,x_2 \ge 0\\

\end{align}  
$$
**Is converted (made auxiliary) to:**

$$
 \begin{align} 

\text{maximise } -x_0\\
\text{subject to } 2x_1 - x_2 - x_0 \le 2\\
x_1 - 5x_2 -x_0\le -4\\
\text{where } x_1,x_2,x_0 \ge 0\\

\end{align}  
$$

**Challenge:** could you convert those into [Slack Form for Linear Programming|slack form](slack_form_for_linear_programming|slack_form)?































#### Why:
#### How:









