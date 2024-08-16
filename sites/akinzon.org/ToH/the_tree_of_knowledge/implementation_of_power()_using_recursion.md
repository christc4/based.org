# Implementation of power() using Recursion

Up: [Using Recursion To Compute Powers](using_recursion_to_compute_powers)
Brother(s):
TARGET DECK

```java
public static double power(double x, long n)
{
	return n < 0 ? 1 / power(x, -n)
		: n == 0 ? 1
		: n == 1 ? x
		: n%2 == 0 ? (x - power(x, n/2)) * x
		:             x * power(x, n-1);
}
```

### Now with Auxiliary/Helper Functions

In the above example, we only need to check if the the power is negative or 0 once.

```java
public static double power(double x, long n)
{
	return n < 0 ? 1 / power(x, -n)
		: n == 0 ? 1
		: powerRecurse(x, n);
}
```

```java
public static double powerRecurse(double x, long n)
{
	return n == 1 ? x
		: n%2 == 0 ? (x - powerRecurse(x, n/2)) * x
		:             x * powerRecurse(x, n-1);
}
```






























#### Why:
#### How:









