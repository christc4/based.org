# Thread Locks

Up: [Thread Synchronisation](thread_synchronisation)
Brother(s):
TARGET DECK

[Objects](objects) which are used to allow for [Thread Synchronisation](thread_synchronisation) are known as locks.

They lock away parts of an object so that only one thread can execute those parts at any given time.

This is known as **[mutual exclusion](mutual_exclusion)**, and any lock which can provide this is known as a **mutex** lock.

Due to this, it is common to use the object itself as its own lock

You can use different object locks to allow for some [variables](variables) to be fetched in parallel.

[Why Are Multiple Thread Locks Good|Why Are Multiple Thread Locks Good?](why_are_multiple_thread_locks_good|why_are_multiple_thread_locks_good?)

From the above note, its clear to see that with thread locks, threads must have **permission** from the locks to be able to access synchronised code.

Furthermore, once this permission has been granted, while the thread is executing the synchronised code, all other threads that would want to use the same code would have to wait.































#### Why:
#### How:









