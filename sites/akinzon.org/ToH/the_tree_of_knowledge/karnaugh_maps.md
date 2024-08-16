# Karnaugh Maps

Up: [Computer Arithmetic](computer_arithmetic)
Brother(s):
TARGET DECK

**Karnaugh Maps** are a convenient way to represent boolean functions with $\leq$ 4 variables 

Elements are arranged in a way where from left to right, or from top to bottom, only one bit change happens. This is known as **Grey code**, and it allows us to track the change of outputs easier.


> [!TIP]- K-Map images:
> **2-variable K-map:**
![Pasted image 20240109172858.png](pasted_image_20240109172858.png)
**3-variable K-map:**
![Pasted image 20240109172915.png](pasted_image_20240109172915.png)
**4-variable K-map:**
![Pasted image 20240109172934.png](pasted_image_20240109172934.png)

[Solving Karnaugh Maps](solving_karnaugh_maps)
### I don't care

don't cares (marked by x's) indicate that we do not care about the output for a particular assignment to boolean variables. 

These can be used to help simplify expressions.

![Pasted image 20240109173646.png](pasted_image_20240109173646.png)

We don't need all don't cares but the one at a'bcd can be used to form an implicant group that helps form a simple expression.

**Implicants** - Terms covering at least one 1 element of a function, and no 0 element

**Prime implicants** - Implicants that cannot be covered by another.

**Essential prime implicants** - Prime implicants that contain at least one 1 element that is not covered by any other implicants.





































#### Why:
#### How:









