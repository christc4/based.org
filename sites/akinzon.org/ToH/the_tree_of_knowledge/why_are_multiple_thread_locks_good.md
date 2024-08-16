# Why Are Multiple Thread Locks Good?

Up: [Thread Locks](thread_locks)
Brother(s):
TARGET DECK

If we had two [variables](variables) `a` and `b` that we wanted to be able to update and access in parallel.

Using `this` as the lock would mean that only one thread would be able to access/modify these variables at the same time.

With a separate lock for `a` and another one for `b`, we would be able to allow one thread to access/modify `a` with its lock,

and the same for `b` as well!

And if we wanted to have a [Methods|method](methods|method) `setBoth()` that could allow both `a` and `b` to be updated at the same time by one thread, we would use **both locks** to lock this code away

**Why?** Let's think about the different cases:

**`setBoth()` has the same lock as `setA()`:** `setBoth()` modifies the contents of `b` with `a`'s lock, so a thread with `a`'s lock and another thread with `b`'s lock can both update `b` at the same time, which would lead to confusion over the correct value!

**`setBoth()` has the same lock as `setB()`:** Same problem as for `setA()`, but instead the value of `a` becomes the confusing one!

**`setBoth()` has its own lock distinct from `setA` and `setB`:** Well, this is just an even worse version of the two above issues, since now **BOTH** `a` **AND** `b` can now be updated at the same time!!! (**JISOS COME DOWN NOW**)

But by using both `a` and `b`'s locks for `setBoth()`, only the thread with both locks can access/modify `a` and `b` during this time.



































#### Why:
#### How:









