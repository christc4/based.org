# Yes, You Can Over-Synchronise (Downside of synchronisation)

Up: [Thread Synchronisation](thread_synchronisation)
Brother(s):
TARGET DECK

Synchronising too much code leads to other threads blocking, waiting for one thread to be finished with the [Thread Locks|lock](thread_locks|lock).

There is also a large overhead to using locks.

That's why it is best to synchronise as little as possible, to minimise the penalty.





































#### Why:
#### How:









