# Notifying Threads

Up: [Communicating Between Threads](communicating_between_threads)
Brother(s):
TARGET DECK

When a [Threads|thread](threads|thread) is done with executing some code from within a [Monitors|monitor](monitors|monitor), it can call `Monitor.notifyAll()` or `Monitor.notify()` to tell all threads in the [Monitor Wait Queue](monitor_wait_queue) that it has finished.

`Monitor.notify()` only works on one random, arbitrary thread.

`Monitor.notifyAll()` wakes up all threads, and can be used/allows for one specific thread to be woken up.

$$ \usepackage{algorithmic} $$
$$ \usepackage{algorithmic}  
\begin{algorithmic} 
1+1=2 \label{eq:adds} 




$$
































#### Why:
#### How:









