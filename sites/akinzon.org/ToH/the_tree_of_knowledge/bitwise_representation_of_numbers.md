# Bitwise Representation of Numbers

Up: [Computer Arithmetic](computer_arithmetic)
Brother(s):
TARGET DECK


## We have...


#### Unsigned integers (ew)

> [!TIP]- Image representation
> ![Pasted image 20240109180002.png](pasted_image_20240109180002.png)


#### Signed integers

##### Signed Magnitude 
In signed magnitude, the MSB acts as the sign bit (0 = +, 1 = -)

>[!TIP]- Image representation
>![Pasted image 20240109180147.png](pasted_image_20240109180147.png)

##### Two's Complement
In Two's complement, $x$ + $-x$ = $2^n$ 

To get to two's complement, flip the bits and add 1.

It's a good choice for hardware since it allows for the use of an unsigned adder (adders that do not need to track the sign of a value)

##### One's Complement
Just flip the bits, don't add 1

##### Biased offset
The whole point of bias is that it sets some value as our new 0.
For all values $V$, $V \geq -B$.

> [!TIP]- Different numerical representations
> ![Pasted image 20240109181016.png](pasted_image_20240109181016.png)


#### Floating Point

>[!TIP]- Image representation
>![Pasted image 20240109181400.png](pasted_image_20240109181400.png)

This is the **IEEE standard: 754-1985**

**For an unsigned binary value**: Normalise it, put it 1.{Rest of value}

For every jump left the decimal place makes, add 1 to exponent, every jump right takes 1 away.

Add the bias to the exponent

It is now stored in the format:

sign|exponent+bias|(implied 1 due to normalisation)|every value after the radix point(fraction field)

Given this value:
![Pasted image 20240109182312.png](pasted_image_20240109182312.png)

Find its IEEE floating point representation

>[!TIP]- Answer:
>Normalise: 1.00111
>Exponent is 5
>Biased exponent: 5 + 127 = 132 (10000100)
>Sign: 0 (positive)
>Fraction field: 00111
>Final answer:
>0 10000100 00111 000000000000

























#### Why:
#### How:









