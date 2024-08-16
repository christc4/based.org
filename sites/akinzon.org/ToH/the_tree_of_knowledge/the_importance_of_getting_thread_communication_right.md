# The Importance of Getting Thread Communication Right

Up: [Communicating Between Threads](communicating_between_threads)
Brother(s):
TARGET DECK

At first, communicating between threads seems easy:
	- Allow threads to modify [objects](objects) in [the heap](the_heap) 
	- Allow them to then access these objects
	- Profit

NUH-UH!! It's not this easy!!!

If some threads try to read data before other threads have finished writing it, outdated data may end up being used, leading to a corruption of the data!

[Thread Synchronisation](thread_synchronisation) doesn't completely solve this issue either, since it only means that threads are stuck waiting for the other thread to give up it's [Thread Locks|lock](thread_locks|lock) before it can access other code. This leads to threads perpetually performing the same code over and over (what a shame...)

We need some way to tell threads **immediately** that they can access the code they were waiting for...




































#### Why:
#### How:









