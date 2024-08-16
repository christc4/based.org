# Use of the p Array for The WIS Algorithm

Up: [The Weighted Interval Scheduling Algorithm](the_weighted_interval_scheduling_algorithm)
Brother(s):
TARGET DECK

We also make sure to sort the intervals so that we do not need to search through **every** interval to check for overlaps (all intervals behind the first interval that does not overlap with $x_n$ do not overlap either.)

Integral to this algorithm is an [Arrays|array](arrays|array) `p` where `p[j] = k` is the largest index `k < j` such that intervals $i_k$ and $i_j$ do not overlap.

Because of this, any interval with index $\le$ `p[j]` cannot have an overlap with $i_j$ either.

The `p` array is op!




































#### Why:
#### How:









