# Radix Sort + Counting Sort

Up: [Radix Sort](radix_sort), [Counting Sort](counting_sort)
Brother(s):
TARGET DECK

![Pasted image 20240306074549.png](pasted_image_20240306074549.png)

The time complexity of this combination is $\Theta(d(n+k))$ since counting sort is $\Theta(n+k)$ and this combo sort runs counting sort $d$ times.

When $d$ is constant, it can be ignored and the complexity defaults to just counting sort.

This combo is [Stability of an Algorithm|stable](stability_of_an_algorithm|stable), but not [When an Algorithm is In-place|in-place](when_an_algorithm_is_in-place|in-place) since the counting sort is not in-place.


































#### Why:
#### How:









