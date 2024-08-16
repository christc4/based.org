# Arrays (Algorithmics)

Up: [List of Some Data Structures](list_of_some_data_structures)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1201 Algorithmics::Data Structures

Arrays use a contiguous (consecutive) chunk of memory, with an access time of $\Theta (1)$. 

They are very memory efficient in this regard, and give a high performance 95% of the time (cap)

Arrays are of a fixed length usually, so adding new items is a pain. (But [Variable-Length Arrays](variable-length_arrays) can be used)

Inserting/deleting anything in the middle is $\Theta (n)$

Why is array access time $\Theta (1)$, but insertion to the middle  $\Theta (n)$? #flashcard 
Array access time is $\Theta (1)$ because it is stored in a chunk of memory, any array index is simply an offset of the first element. Insertion is $\Theta (n)$ because you have to shift all other values to the right, which asymptotically grows at a linear rate.
<!--ID: 1714246190460-->



































#### Why:
#### How:









