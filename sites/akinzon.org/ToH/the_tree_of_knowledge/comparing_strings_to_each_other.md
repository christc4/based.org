# Comparing Strings to Each Other

Up: [Strings](strings)
Brother(s):
TARGET DECK

in Python (heh, python is mid), you might've seen this:
```python
if string == "hello":
	# code
```

..Yeah that doesn't have the same effect in Java. Using `"=="` in Java tests for **identity,** not **equality.** So:

```java
if (input == otherInput) {
	// do smth;
}
```

Will test if two values are the exact same variable.

The java-save way of testing for equality between strings is:

```java
if (input.equals(otherInput)) {
	// do smth;
}
```

![Pasted image 20240202135828.png](pasted_image_20240202135828.png)
double equals will only evaluate to true for `str2 == str3`, since they both point to 0x1002.

[Advanced String Comparison](advanced_string_comparison)































#### Why:
#### How:









