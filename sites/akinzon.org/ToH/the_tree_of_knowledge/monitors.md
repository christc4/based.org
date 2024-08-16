# Monitors

Up: [Communicating Between Threads](communicating_between_threads)
Brother(s):
TARGET DECK

Monitors are an abstraction of the way to think about controlling concurrency.

The monitor is set to be the object used as the thread lock.

Monitors consist of:
	- Code fragments protected by mutex [Thread Locks|locks](thread_locks|locks)
	- Threads that make themselves wait in the monitor

[How Threads Work With Monitors](how_threads_work_with_monitors)

Any [Objects|object](objects|object) can be a monitor, since they can be used as their own mutex lock and also has a [Monitor Wait Queue|queue](monitor_wait_queue|queue) for threads waiting to access its code!

Only code wrapped in [The Synchronized Block](the_synchronized_block) with lock `o` will be part of `o`'s monitor.


































#### Why:
#### How:









