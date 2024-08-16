# Hoare Scheme for Finding the Partition Point

Up: [Quicksort](quicksort)
Brother(s):
TARGET DECK

Hoare's scheme starts with two pointers, `i` and `j`, which increment and decrement respectively until `i` reaches a value greater than the pointer and `j` reaches a value less than the pointer. 

If `i` is less than `j` in terms of index at this point, then the values at `i` and `j` get swapped. Else the value of `i` is returned as the partition point.

At this point, all values left of `i` are $\le$ the pivot, and all values to the right are greater.

But that's it. Unlike the [Lomuto Scheme for Finding the Partition Point|Lomuto Scheme](lomuto_scheme_for_finding_the_partition_point|lomuto_scheme), the pivot itself isn't placed in the correct position, and still has to be sorted later.

We essentially partition the list into values smaller than (or equal to) the pivot and values greater than (or equal to) but the original pivot itself isn't placed anywhere specific in this scheme.

![Pasted image 20240305201808.png](pasted_image_20240305201808.png)































#### Why:
#### How:









