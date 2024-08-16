# Recursive note

	package main

	import "fmt"

	func fact(n int) int {
		if n == 0 {
			return 1
		}
		return n * fact(n-1)
	}

	func main() {
		fmt.Println(fact(7))

		var fib func(n int) int
		
		fib = func(n int) int {
			if n < 2 {
				return n
			}

			return fib(n-1) + fib(n-2)
		}

		fmt.Println(fib(7))
	}

## factorial function `fact`

- start with _n_ = 7
	- `fact(7) = 7 * fact(6)`

- do `fact(6)`
	- `fact(6) = 6 * fact(5)`

- and so on…
	- `fact(1) = 1 * fact(0)`

- base case
	- `fact(0) = 1`

## get calculator

calculate each step from bottom up

- fact(0) = 1
- fact(1) = 1 * 1 = 1
- fact(2) = 2 * 1 = 2
- fact(3) = 3 * 2 = 6
- fact(4) = 4 * 6 = 24
- fact(5) = 5 * 24 = 120
- fact(6) = 6 * 120 = 720
- fact(7) = 7 * 720 = 5040

So, fact(7) is 5040.

## fibonacci function

> fib sequence adds two previous no. to get the next number

- start with _n_ = 7
	`fib(7) = fib(6) + fib(5)`

- compute `fib(6)`
	`fib(6) = fib(5) + fib(4)`

- and so on
	- base cases are `fib(1) = 1` &amp; `fib(0) = 0`

## get calculator

- fib(0) = 0
- fib(1) = 1
- fib(2) = 1 + 0 = 1
- fib(3) = 1 + 1 = 2
- fib(4) = 2 + 1 = 3
- fib(5) = 3 + 2 = 5
- fib(6) = 5 + 3 = 8
- fib(7) = 8 + 5 = 13

### conclusion

after running go build on the source code above, your ouptut should be

5040
13

## no understand

why 0? because `if n == 0` 

what is `return 1`

this is the base case for recursive function, base cases are essential in any recursive function to stop the recursion, otherwise it'd call indefititely, causing a stack overflow

why 1? 

this is a math conveention

if n < 2 {
    return n
}
This handles two base cases for the Fibonacci function:

When 
𝑛
n is 0:
fib(0) should return 0 because the 0th Fibonacci number is 0.
When 
𝑛
n is 1:
fib(1) should return 1 because the 1st Fibonacci number is 1.
Why return n?:
The Fibonacci sequence is defined as:

fib
(
0
)
=
0
fib(0)=0
fib
(
1
)
=
1
fib(1)=1
For 
𝑛
≥
2
n≥2, 
fib
(
𝑛
)
=
fib
(
𝑛
−
1
)
+
fib
(
𝑛
−
2
)
fib(n)=fib(n−1)+fib(n−2)
By returning n when 
𝑛
n is less than 2, we correctly handle the first two numbers of the sequence, allowing the recursive calls to work correctly for larger values of 
𝑛
n.

Summary
return 1 in fact: This is the base case for the factorial function. It ensures that when 
𝑛
n is 0, the function returns 1, stopping further recursion. This is because the factorial of 0 is defined to be 1.

return n in fib: This handles the base cases for the Fibonacci function. It ensures that when 
𝑛
n is 0 or 1, the function returns 
𝑛
n (0 or 1, respectively), stopping further recursion. This is because the first two Fibonacci numbers are defined as 0 and 1.

base case breaks the chain of recursion, it is sometimes also called the "terminating case".
Neglecting to write a base case, or testing for it incorrectly, can cause an infinite loop.
## Now go read...

[wikipedia](https://en.wikipedia.org/wiki/Recursion_(computer_science)
