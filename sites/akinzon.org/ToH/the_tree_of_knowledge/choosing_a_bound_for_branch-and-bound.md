# Choosing a Bound for Branch-and-Bound

Up: [Branch-And-Bound](branch-and-bound)
Brother(s):
TARGET DECK

The bound used for branch-and-bound depends on what is trying to be accomplished.

My Akin guess is that if you're trying to maximise a value, you use an upper bound and backtrack on branches that will not get you closer to that upper bound than you can get. ([Branch-And-Bound 0-1 Knapsack](branch-and-bound_0-1_knapsack))

When trying to minimise a value, you would do the same but with a lower bound and backtrack on branches that will not get you closer to the lower bound. ([Branch-And-Bound Integer Linear Programming](branch-and-bound_integer_linear_programming))

When trying to get an exact value, you would use both an upper and lower bound to stop you from going too far in either direction. ([Branch-And-Bound Subset Sum](branch-and-bound_subset_sum))

> [!info] The bound is unique to each problem, so get thinking!

































#### Why:
#### How:









