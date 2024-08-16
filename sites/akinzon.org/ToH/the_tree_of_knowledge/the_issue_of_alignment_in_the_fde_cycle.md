# The issue of Alignment in the [Instruction cycle|FDE cycle](instruction_cycle|fde_cycle)

Up: [Instruction Sets](instruction_sets)
Brother:
TARGET DECK

Memory is addressed in bytes, but is organised in words (32/64 bit),
and the issue is user-defined objects can be of a size that does not fit snugly into either, and so is *misaligned.* Some of this data may be in different blocks, so putting it all together would require multiple memory reads to each block, and shifting the data so that it is as it once was.







































#### Why:
#### How:









