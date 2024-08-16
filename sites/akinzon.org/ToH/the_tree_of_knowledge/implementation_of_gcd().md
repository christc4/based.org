# Implementation of gcd()

Up: [Euclid's GCD with Recursion](euclid's_gcd_with_recursion)
Brother(s):
TARGET DECK

Too ez

```java
public static long gcd(long a, long b)
{
	if (b==0)
		return a;
	else
		return gcd(b, a%b);
}
```



































#### Why:
#### How:









