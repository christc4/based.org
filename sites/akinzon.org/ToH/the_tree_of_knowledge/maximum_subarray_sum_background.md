# Maximum Subarray Sum Background

Up: [The Maximum Subarray Sum](the_maximum_subarray_sum)
Brother(s):
TARGET DECK

Given an [Arrays|array](arrays|array)/[Sets (ADT)|set](sets_(adt)|set) $A$ of numbers (with at least one negative number to stop the solution from being trivially easy), the maximum subarray sum is the subarray $B$ (of **contiguous** elements) whose sum of elements is the largest possible.

We are looking for:

$$ max_{ij} \biggl ( S_{ij} = \sum_{k=i}^{j \ge i } A_k \biggr) $$


We could define a maximal partial suffix $max_i\ S_{ij}$, to make the following true:

$$ max_{ij}\ S_{ij} = max_j\ M_j $$
Since we assume that we have the maximal $i$.

The next step is to calculate the maximal $M_j$ for all $j$.





























#### Why:
#### How:









