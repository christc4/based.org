# Counting Rising Numbers

Up: [Counting via Bijection](counting_via_bijection)
Brother(s):
TARGET DECK

Rising numbers are defined as numbers where the consequent digits are only allowed to ascend.

12 is a rising number, but 132 is not.

> [!TIP]- How many 4-digit rising numbers with $d_4 = 9$ are there?
> Remember that $d_4$ is fixed as 9, so we don't worry about it.
> As a result, we only need to count how many 3-digit rising numbers can be constructed from $\{1, 2, ..., 8\}$
> This is $\binom{8}{3}$

> [!TIP]- With alphabet $\{a, b, c, ..., z\}$, how many 5-letter rising strings with $d_2 = d$ and $d_3 = f$ are there?
> For $d_1$, we have 3 choices $[\binom{3}{1}]$
> $d_2$ and $d_3$ are fixed, so we don't worry about them
> For $d_4$ and $d_5$ we can choose from the remaining alphabet $\{g,...,z\}$ (20 letters) for the next two numbers
> Which is $\binom{20}{2}$  
> The final answer is $\binom{3}{1}\binom{20}{2} = 570$

































#### Why:
#### How:









