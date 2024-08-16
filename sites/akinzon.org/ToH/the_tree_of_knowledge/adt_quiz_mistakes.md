# ADT Quiz Mistakes

Up: [Abstract Data Types](abstract_data_types)
Brother(s):
TARGET DECK

![Pasted image 20240304200206.png](pasted_image_20240304200206.png)

If it was going to be the `N/2` answer, it would have had to be `top1 + 1 != N/2`, because by the time `top1` = `N/2`, it would be too late and a value in the 2nd stack would be overwritten.

****

![Pasted image 20240304200432.png](pasted_image_20240304200432.png)

Ok then, fair enough. Dequeue is as if it were a normal queue, but `deleteMinimum()` technically *is* more catered towards a priority queue. Was a trick question, I guess?

****

![Pasted image 20240304200744.png](pasted_image_20240304200744.png)

Depending on where the elements are in the linked list, there is indeed no guarantee that the element returned is the first or the last of its kind in the list. For example, if both instances are located towards the end of the list, the last one will be returned. If both instances are located towards the start, the first one will be returned. It's not deterministic, and therefore we cannot say what the method will always return.

****

![Pasted image 20240304211730.png](pasted_image_20240304211730.png)

The worst case is also [Big Omega Notation|$\omega$](big_omega_notation|$\omega$)(1), since it acts as a theoretical asymptotic lower bound for the program.

****

![Pasted image 20240304211951.png](pasted_image_20240304211951.png)

For this question, I can't see properly which ones you got right/wrong, but it should be clear now that for [stacks](stacks), all operations take place in $\Theta(1)$ time, so the [Big O Notation|big O](big_o_notation|big_o) and [Big Omega Notation|big omega](big_omega_notation|big_omega) should be 1, and due to the nature of big O, it is also O(n). 
























#### Why:
#### How:









