# Huffman Coding

Up: [Examples of Greedy Algorithms](examples_of_greedy_algorithms)
Brother(s):
TARGET DECK

[Huffman Coding Problem Background](huffman_coding_problem_background)

[Prefix-free Codes](prefix-free_codes)

Huffman coding works by building a prefix-free code binary tree for an alphabet from the ground up, starting with the most frequent symbols and working your way to the least frequent.

The plan is to give the most frequent symbols the smallest codes, and the least frequent symbols the longest codes.

To build an optimal [Trees|tree](trees|tree) $T$ for $k$ symbols:
	- First build the optimal tree $T'$ for $k-1$ symbols, where the two nodes with the smallest frequencies are replaced by a single node with the sum of their frequencies.

For any node in this tree with only once child, it would also be better to replace it with that child
































#### Why:
#### How:









