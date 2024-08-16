# Matrices

Up: [Linear Algebra](linear_algebra)
Brother(s):
TARGET DECK

**Matrices** are a set of numbers arranged in a square/rectangular form
	![Pasted image 20231208201237.png](pasted_image_20231208201237.png)
	d1 - columns
	d2 - rows

Matrices are [Linear Maps|linear maps](linear_maps|linear_maps): They can turn one vector into another.
	![Pasted image 20231208201326.png](pasted_image_20231208201326.png)


**Index form for applying a matrix transformation**:
	$$ (Av)_i = \sum_{j=1}^{d_1} a_{ij}v_j$$

i - row
j - column

$$ (Av)_2=∑_{j=1} ^2 a_{2j} v_j=a_{21} v_1+a_{22} v_2=3∗1+1∗2=5 $$

#### Matrix multiplication:
![WhatsApp Image 2024-01-08 at 18.18.46_79156ccb.jpg](whatsapp_image_2024-01-08_at_18.18.46_79156ccb.jpg)
Is a big for loop:
$$ C_{kj}=∑_i A_{ki} B_{ij} $$
Matrix multiplication is a matrix transformation on $n$ vectors, where $n$ is the number of columns in the rightmost vector.

You're effectively changing where the basis vectors which represent matrix $B$ land.

**Rules:**

Matrix multiplication is distributive: $$A(αB+βC)=α AB+β AC , α,β∈R$$
Matrix multiplication is associative:
$$ A(BC)=( AB)C $$

Matrix multiplication is generally non-commutative:
$$ AB≠BA $$


[Matrix Transformations in Different Coordinate Systems](matrix_transformations_in_different_coordinate_systems)

[Inverse Matrices](inverse_matrices)

[The Determinant](the_determinant)




















#### Why:
#### How:









