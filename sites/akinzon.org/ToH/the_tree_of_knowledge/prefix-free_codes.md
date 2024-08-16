# Prefix-free Codes

Up: [Huffman Coding](huffman_coding)
Brother(s):
TARGET DECK

Prefix-free codes are symbol encodings where no code is a prefix of another.

The allow us to use shorter codes for more frequent symbols and longer ones for less frequent symbols because without prefix-free codes, it would be ambiguous what a code was meant to represent.

For example if `a` was represented by 0, and `b` was represented by `01`, it would be ambiguous for a parser since it wouldn't know if the correct value was `b` or to parse it as `0` = `a` followed by a `1`.

Prefix-free codes remove all that ambiguity!

Prefix-free codes can be represented by [Binary Trees](binary_trees).


































#### Why:
#### How:









