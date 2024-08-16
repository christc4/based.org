# Algorithm Analysis Questions

Up: [Algorithm Analysis](algorithm_analysis)
Brother(s):
TARGET DECK
Leaves of the Tree of Ohara::Branch I::Semester II::1201 Algorithmics::Algorithm Analysis



What does it mean for a sorting algorithm to be [Stability of an Algorithm|stable](stability_of_an_algorithm|stable)? #flashcard 
Stable algorithms do not change internal order when sorting data. 
Data that already has some order to it has its order preserved.
<!--ID: 1709673394271-->





What does it mean for a sorting algorithm to be [When an Algorithm is In-place|in-place](when_an_algorithm_is_in-place|in-place)? #flashcard 
In-place algorithms have constant memory usage and do not require additional $n$-sized memory to work.
<!--ID: 1709673394279-->





Compare [insertion sort](insertion_sort), [selection sort](selection_sort) and [bubble sort](bubble_sort) based on running time (in all 3 cases), and how [Stability of an Algorithm|stable](stability_of_an_algorithm|stable) and [When an Algorithm is In-place|in-place](when_an_algorithm_is_in-place|in-place) they are. #flashcard 
All 3 algorithms are $\Theta(n^2)$ in the average and worst case, with selection sort being the only one of the three that has an $\Theta(n^2)$ best case due to it still running $n^2$ times even in the event of a sorted input.
All three algorithms are in-place, but selection sort is unstable since it can mess with the relative order of equal elements.
<!--ID: 1709673394288-->





When might [insertion sort](insertion_sort) be beneficial? When might [selection sort](selection_sort) be beneficial? What are the pros and cons of [bubble sort](bubble_sort)? #flashcard 
Insertion sort is very fast on small lists.
Selection sort saves some swaps and can be very good on in places where swapping is expensive.
Bubble sort is easy to implement but **no-one** uses it in practice (LOL)
<!--ID: 1709673394294-->





What are the advantages of [merge sort](merge_sort) over [quicksort](quicksort) and vice versa? #flashcard 
Merge sort is $\Theta(nlog(n))$ in all cases (asymptotically optimal) whereas quicksort is $\Theta(n^2)$ in the worst case (DAMN). Merge sort is also [Stability of an Algorithm|stable](stability_of_an_algorithm|stable) unlike quicksort.
Merge sort makes fewer comparisons.
Quicksort has less data movement.
Quicksort is slightly faster and is [When an Algorithm is In-place|in-place](when_an_algorithm_is_in-place|in-place), unlike merge sort.
<!--ID: 1709673394302-->





Evaluate [Counting Sort](counting_sort) based on [Stability of an Algorithm|stability](stability_of_an_algorithm|stability) and [When an Algorithm is In-place|in-place](when_an_algorithm_is_in-place|in-place)ness, as well as [Running Time of an Algorithm|running time](running_time_of_an_algorithm|running_time). Do the same for a [Radix Sort](radix_sort) which uses a counting sort. #flashcard 
Counting sort is stable, but it is not in-place since it requires n-sized memory to store the sorted array. Counting sort is $\Theta (n + k)$. Radix sort would also be stable and not in-place, but its running time would be $\Theta (d(n + k))$ since it runs the counting sort $d$ times.
<!--ID: 1709755705850-->




While [Radix Sort + Counting Sort](radix_sort_+_counting_sort) may be $\Theta (d(n + k))$ in theory, why is it not the GOAT of all [sorting algorithms](sorting_algorithms) then? #flashcard 
While counting sort is the theoretical GOAT, in practice it would require additional operations to split up the digits of each element in the array, and also additional memory, and with all the data transfer going on, the hardware becomes a bottleneck for this [Sorting Algorithms|sorting algorithm](sorting_algorithms|sorting_algorithm).
<!--ID: 1709755705860-->
























#### Why:
#### How:









