# Thread Synchronisation

Up: [Control Flow](control_flow)
Brother(s):
TARGET DECK

Multiple threads executing the same code at the same time can be tricky: [Problems with Shared Variable Access For Threads](problems_with_shared_variable_access_for_threads)

This is where **synchronisation** comes in!

**Synchronisation** of operations allows for operations to be **locked** when a thread is executing them, such that other threads are forced to wait their turn, regardless of the scheduler's order.

Only one thread can get permission to execute thread-safe code at a time.

It somewhat forces a sequential order (cool, eh?)

[Using Synchronisation in Java](using_synchronisation_in_java)

[Thread Locks](thread_locks)

[Mutual Exclusion](mutual_exclusion)

[Atomic Blocks](atomic_blocks)

[Yes, You Can Over-Synchronise](yes,_you_can_over-synchronise)

[Synchronisation Rules of Thumb](synchronisation_rules_of_thumb)

[Thread Deadlocks](thread_deadlocks)























#### Why:
#### How:









