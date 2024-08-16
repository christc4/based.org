# Monitor Wait Queue

Up: [Monitors](monitors)
Brother(s):
TARGET DECK

This [Queues|queue](queues|queue) is used for [threads](threads) that have already executed the code from inside a monitor and are placing themselves in the queue to be able to access the monitor again.

They are placed in the queue by [Making a Thread Wait|waiting](making_a_thread_wait|waiting), calling `Monitor.wait()`

They can be released from the queue by being [Notifying Threads|notified](notifying_threads|notified), when `Monitor.notify()` or `Monitor.notifyAll()` is called.

A monitor can essentially have multiple wait queues by separating queues using flag [variables](variables) to cause certain threads to wait for certain flags.

































#### Why:
#### How:









