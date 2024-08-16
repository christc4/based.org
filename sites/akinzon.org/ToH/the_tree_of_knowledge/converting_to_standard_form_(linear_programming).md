# Converting to Standard Form (Linear Programming)

Up: [Standard Form for Linear Programming](standard_form_for_linear_programming)
Brother(s):
TARGET DECK

These are the steps for converting to standard form:
	- Any minimisation problems $min\ f$ are converted to $max\ -f$.
	- For any [Linear Programming Contstraints|constraint](linear_programming_contstraints|constraint):
		- $\sum_{j=1}^n a_{ij}x_j = b$ is replaced with $\sum_{j=1}^n a_{ij}x_j \le b$ and $\sum_{j=1}^n a_{ij}x_j \ge b$.
		- $\sum_{j=1}^n a_{ij}x_j \ge b$ is replaced with $\sum_{j=1}^n -a_{ij}x_j \le -b$.
		- If a variable $x_j$ has no sign restriction (what is a sign restriction?) it is replaced by $x_j = x_j^+ - x_j^-$, where $x_j^+,x_j^- \ge 0$.


Sign restrictions are positivity/negativity constraints:
![Pasted image 20240503170917.png](pasted_image_20240503170917.png)

































#### Why:
#### How:









