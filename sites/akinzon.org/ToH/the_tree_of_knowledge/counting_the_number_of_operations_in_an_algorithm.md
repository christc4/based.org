# Counting The Number of Operations in an Algorithm

Up: [Asymptotic Notation](asymptotic_notation)
Brother(s):
TARGET DECK

To get a good grasp for the number of operations, count the number of [For Loop](for_loop)s and nested for loops, for example:
![Pasted image 20240130210003.png](pasted_image_20240130210003.png)
This will be a $\Theta (n)$ algorithm as it has a single for loop, and
![Pasted image 20240130210031.png](pasted_image_20240130210031.png)
this will be a $\Theta (n^2)$ algorithm due to its nested for loop structure.

With this however,
![Pasted image 20240130210130.png](pasted_image_20240130210130.png)
Its a bit trickier to say what the average will be, because the for loop locked away by the if statement may not always run, so the average is somewhere between $\Theta (n)$ and $\Theta (n^2)$.

[Bounds For the Running Time](bounds_for_the_running_time)

































#### Why:
#### How:









