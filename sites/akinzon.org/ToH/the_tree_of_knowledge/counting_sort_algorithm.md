# Counting Sort Algorithm

Up: [Counting Sort](counting_sort)
Brother(s):
TARGET DECK


![Pasted image 20240306074218.png](pasted_image_20240306074218.png)

This algorithm takes the unsorted array, an empty array for the result and a value $k$ as input.

This algorithm first creates a counting array of size $k$, and initialises it to be full of 0's.

It then counts how many occurrences there are of each value in the unsorted array.

Each value in the counter array is stored in the index that is the same value as the value in the original unsorted array. For example, if 3 was in the original array, its frequency would be stored at index 3 in the counter array.

It then cumulatively counts these occurrences so that the last item in the counter array is the total number of items, and is the same as the last index in the original array.

These cumulative counts become the array indices for the values in the sorted array

The original array is then iterated through backwards, and each value's cumulative count is used as its index in the sorted array and then decremented, so that all duplicate values can be stored next to each other.





































#### Why:
#### How:









