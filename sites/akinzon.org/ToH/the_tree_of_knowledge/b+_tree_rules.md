# B+ Tree Rules

Up: [B+ Trees](b+_trees)
Brother(s):
TARGET DECK

**Rules:**
	1) Actual data is only stored in leaves
	2) Non-leaf nodes store $M-1$ keys to aid navigation
	3) The root is either a leaf or has between 2 and $M$ children
	4) All non-leaf nodes have between $\lceil M/2 \rceil$ and $M$ children
	5) All leaves (except the root) are at the same level and have between $\lceil L/2 \rceil$ and $L$ data entries

How $M$ and $L$ are chosen depends on how much is being read/written at a time.



































#### Why:
#### How:









