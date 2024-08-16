# How [Threads](threads) Work With [Monitors](monitors)

Up: 
Brother(s):
TARGET DECK

Threads will make themselves wait by calling `Monitor.wait()`, essentially adding itself to the [Monitor Wait Queue|monitor's wait queue](monitor_wait_queue|monitor's_wait_queue).

Calling `Monitor.notify()` allows for the first thread in the queue to be [Notifying Threads|notified](notifying_threads|notified) that it can now access the critical code fragment because the lock has been released (So long as it actually has).



































#### Why:
#### How:









